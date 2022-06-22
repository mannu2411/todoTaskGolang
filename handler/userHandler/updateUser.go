package userHandler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/todoTask/database/helper"
	"github.com/todoTask/models"
)

func UpdateRow(writer http.ResponseWriter, request *http.Request) {
	var update_user models.UpdateUser
	decoder := json.NewDecoder(request.Body)
	addErr := decoder.Decode(&update_user)
	log.Printf(update_user.Name)
	if addErr != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	userID, err := helper.UpdateUser(update_user.Name, update_user.Email, update_user.ID)
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
