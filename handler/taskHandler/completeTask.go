package taskHandler

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/todoTask/database/helper"
	"github.com/todoTask/models"
)

func CompleteTask(writer http.ResponseWriter, request *http.Request) {

	/* _, err, flag := utilities.MiddlewareAuth(writer, request)

	if err != nil {
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}
	if flag {
		writer.WriteHeader(http.StatusUnauthorized)
		return
	} */

	var completeTask models.Task

	addErr := json.NewDecoder(request.Body).Decode(&completeTask)
	//log.Printf(completeTask.ID)
	if addErr != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	task, err := helper.CompleteTask(completeTask.ID)
	log.Printf(task)
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
	session := request.Header.Values("session_token")
	sessionId := session[0]
	err = helper.RefreshSession(expireAt, sessionId)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	}
	writer.Write(jsonData)
}
