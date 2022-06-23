package taskHandler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/todoTask/database/helper"
)

func GetTasks(writer http.ResponseWriter, request *http.Request) {
	/* sessionId, err, flag := utilities.MiddlewareAuth(writer, request)
	log.Printf(sessionId)
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
	log.Printf(email)
	if err != nil || email == "" {
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	task, err := helper.GetAllTasks(email)
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
