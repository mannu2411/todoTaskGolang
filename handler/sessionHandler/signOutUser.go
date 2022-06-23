package sessionHandler

import (
	"net/http"

	"github.com/todoTask/database/helper"
)

func SignOut(writer http.ResponseWriter, request *http.Request) {
	/* sessionId, err, flag := utilities.MiddlewareAuth(writer, request)
	//log.Printf(sessionId)
	if err != nil {
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}
	if flag {
		writer.WriteHeader(http.StatusUnauthorized)
		return
	} */
	session := request.Header.Values("session_token")
	sessionId := session[0]
	err := helper.DeleteSession(sessionId)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}
