package userHandler

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/todoTask/database/helper"
	"github.com/todoTask/models"
)

func UpdateUser(writer http.ResponseWriter, request *http.Request) {
	var updateUser models.UpdateUser
	addErr := json.NewDecoder(request.Body).Decode(&updateUser)
	if addErr != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	userID, err := helper.UpdateUser(updateUser.Name, updateUser.Email, updateUser.ID)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Printf("User: " + userID + " has been updated.")
	jsonData, jsonErr := json.Marshal(updateUser)
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
