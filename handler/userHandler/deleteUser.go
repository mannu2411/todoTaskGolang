package userHandler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/todoTask/database/helper"
	"github.com/todoTask/models"
)

func DeleteRow(writer http.ResponseWriter, request *http.Request) {
	var delete_user models.UpdateUser
	addErr := json.NewDecoder(request.Body).Decode(&delete_user)
	log.Printf(delete_user.ID)
	if addErr != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	userID, err := helper.DeleteUser(delete_user.ID)
	log.Printf(userID)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	user, userErr := helper.GetUser(userID)
	if userErr != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	jsonData, jsonErr := json.Marshal(user)
	if jsonErr != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.Write(jsonData)
}
