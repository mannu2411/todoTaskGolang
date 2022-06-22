package taskHandler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/todoTask/database/helper"
	"github.com/todoTask/utilities"
)

func GetTasks(writer http.ResponseWriter, request *http.Request) {
	sessionId, err, flag := utilities.MiddlewareAuth(writer, request)
	log.Printf(sessionId)
	if err != nil {
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}
	if flag {
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}
	task, err := helper.GetAllTasks(sessionId)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	jsonData, jsonErr := json.Marshal(task)
	if jsonErr != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.Write(jsonData)
}
