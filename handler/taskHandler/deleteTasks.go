package taskHandler

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/todoTask/database/helper"
	"github.com/todoTask/models"
)

func DeleteTask(writer http.ResponseWriter, request *http.Request) {

	/* _, err, flag := utilities.MiddlewareAuth(writer, request)
	//log.Printf(sessionId)
	if err != nil {
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}
	if flag {
		writer.WriteHeader(http.StatusUnauthorized)
		return
	} */

	var deleteTask models.Task

	addErr := json.NewDecoder(request.Body).Decode(&deleteTask)
	log.Printf(deleteTask.ID)
	if addErr != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	task, err := helper.DeleteTask(deleteTask.ID)
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
