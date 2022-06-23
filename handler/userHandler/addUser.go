package userHandler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/todoTask/database/helper"
	"github.com/todoTask/models"
)

func AddUser(writer http.ResponseWriter, request *http.Request) {
	var addUser models.AddUser
	addErr := json.NewDecoder(request.Body).Decode(&addUser)
	if addErr != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	userID, err := helper.CreateUser(addUser.Name, addUser.Email, addUser.Pass)

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Printf("User: " + userID + "has been added")
	jsonData, jsonErr := json.Marshal(addUser)
	if jsonErr != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	}
	writer.Write(jsonData)
}
