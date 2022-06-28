package middlewareUtils

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"log"
	"net/http"
	"time"
)

const secretkey string = "SuperSecretKey"

func GenerateJWT(email string) (string, error) {
	var mySigningKey = []byte(secretkey)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}

func JWTAuthMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Token"] == nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var mySigningKey = []byte(secretkey)

		token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("there was an error in parsing")
			}
			return mySigningKey, nil
		})

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if token.Valid {
			log.Printf("JWT Token verified...")
			handler.ServeHTTP(w, r)
		}
		w.WriteHeader(http.StatusUnauthorized)
		return
	})
}
