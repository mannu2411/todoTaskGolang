package middlewareUtils

import (
	"log"
	"net/http"

	"github.com/todoTask/database/helper"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		session := request.Header.Values("session_token")
		sessionId := session[0]

		exist, err := helper.SessionExist(sessionId)
		if err != nil {
			writer.WriteHeader(http.StatusUnauthorized)
			return
		}
		if !exist {
			writer.WriteHeader(http.StatusUnauthorized)
			return
		}
		isExpired, err := helper.IsExpired(sessionId)

		if isExpired {
			writer.WriteHeader(http.StatusUnauthorized)
			return
		}
		log.Printf("Session: " + sessionId + " has verified.")
		next.ServeHTTP(writer, request)
	})
}
