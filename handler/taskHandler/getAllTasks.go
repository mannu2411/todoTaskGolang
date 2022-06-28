package taskHandler

import (
	"encoding/json"
	"github.com/todoTask/middlewareUtils"
	"net/http"
	"time"

	"github.com/todoTask/database/helper"
)

func GetTasks(writer http.ResponseWriter, request *http.Request) {
	session := request.Header.Values("session_token")
	sessionId := session[0]
	/*email, err := helper.GetCreds(sessionId)
	log.Printf(email)
	if err != nil || email == "" {
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}
	*/
	userID := middlewareUtils.UserFromContext(request.Context()).ID
	task, err := helper.GetAllTasks(userID)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	jsonData, jsonErr := json.Marshal(task)
	if jsonErr != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	expireAt := time.Now().Add(360 * time.Second)
	err = helper.RefreshSession(expireAt, sessionId)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	}
	writer.Write(jsonData)
}
