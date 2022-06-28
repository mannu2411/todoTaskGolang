package taskHandler

import (
	"encoding/json"
	"fmt"
	"github.com/todoTask/middlewareUtils"
	"log"
	"net/http"
	"time"

	"github.com/todoTask/database/helper"
	"github.com/todoTask/models"
)

func AddTask(writer http.ResponseWriter, request *http.Request) {

	/*	userID, err := helper.GetCreds(sessionId)

		if err != nil {
			writer.WriteHeader(http.StatusUnauthorized)
			return
		}*/
	userID := middlewareUtils.UserFromContext(request.Context()).ID
	var addTask models.AddTask
	err := json.NewDecoder(request.Body).Decode(&addTask)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	}
	taskID, err := helper.CreateTask(userID, addTask.Task)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	}
	log.Printf(taskID)
	expireAt := time.Now().Add(360 * time.Second)

	session := request.Header.Values("session_token")
	sessionId := session[0]
	err = helper.RefreshSession(expireAt, sessionId)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	}
	writer.Write([]byte(fmt.Sprintf("Task: %s has been created", taskID)))
}
