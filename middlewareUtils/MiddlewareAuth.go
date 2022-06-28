package middlewareUtils

import (
	"context"
	"github.com/todoTask/models"
	"log"
	"net/http"

	"github.com/todoTask/database/helper"
)

const ContextUserKey string = "user"

func UserFromContext(ctx context.Context) models.User {
	return ctx.Value(ContextUserKey).(models.User)
}

/*type HandlerMiddleware interface {
	HandleHTTPC(ctx context.Context, rw http.ResponseWriter, req *http.Request, next http.Handler)
}*/

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		session := request.Header.Values("session_token")
		sessionId := session[0]
		notExist, err := helper.SessionExist(sessionId)
		if err != nil {
			writer.WriteHeader(http.StatusUnauthorized)
			return
		}
		if notExist {
			writer.WriteHeader(http.StatusUnauthorized)
			return
		}
		log.Printf("Session: " + sessionId + " has verified.")
		next.ServeHTTP(writer, request)
	})
}
func GetUserContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session := r.Header.Values("session_token")
		if session == nil {
			next.ServeHTTP(w, r)
		}
		log.Println(r.Method, "-", r.RequestURI)

		sessionId := session[0]
		userID, err := helper.GetCreds(sessionId)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		} else {
			log.Printf("UserID: " + userID)
			users, err := helper.GetUserDetails(userID)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			ctx := context.WithValue(r.Context(), ContextUserKey, users)
			//ctx := context.WithValue(r.Context(), "UserID", userID)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
	})
}
