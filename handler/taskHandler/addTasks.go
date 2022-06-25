package taskHandler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/todoTask/database/helper"
	"github.com/todoTask/models"
)

func AddTask(writer http.ResponseWriter, request *http.Request) {
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
	email, err := helper.GetCreds(sessionId)

	if err != nil {
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}
	var addTask models.AddTask
	err = json.NewDecoder(request.Body).Decode(&addTask)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	}
	taskID, err := helper.CreateTask(email, addTask.Task)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	}
	expireAt := time.Now().Add(360 * time.Second)
	err = helper.RefreshSession(expireAt, sessionId)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	}
	writer.Write([]byte(fmt.Sprintf("Task: %s has been created", taskID)))
}
