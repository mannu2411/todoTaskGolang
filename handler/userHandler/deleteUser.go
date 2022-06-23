package userHandler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/todoTask/database/helper"
	"github.com/todoTask/models"
)

func DeleteUser(writer http.ResponseWriter, request *http.Request) {

	var deleteUser models.UpdateUser
	addErr := json.NewDecoder(request.Body).Decode(&deleteUser)
	if addErr != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	session := request.Header.Values("session_token")
	sessionId := session[0]
	err := helper.DeleteSession(sessionId)
	userID, err := helper.DeleteUser(deleteUser.ID)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Printf("User: " + userID + "has been deleted")
	jsonData, jsonErr := json.Marshal(deleteUser)
	if jsonErr != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.Write(jsonData)
}
