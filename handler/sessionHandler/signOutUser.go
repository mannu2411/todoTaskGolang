package sessionHandler

import (
	"net/http"

	"github.com/todoTask/database/helper"
)

func SignOut(writer http.ResponseWriter, request *http.Request) {
	session := request.Header.Values("session_token")
	sessionId := session[0]
	err := helper.DeleteSession(sessionId)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}
