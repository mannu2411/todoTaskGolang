package userHandler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/todoTask/database/helper"
	"github.com/todoTask/models"
)

func AddRow(writer http.ResponseWriter, request *http.Request) {
	var addUser models.AddUser
	addErr := json.NewDecoder(request.Body).Decode(&addUser)
	log.Printf(addUser.Name)
	if addErr != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	userID, err := helper.CreateUser(addUser.Name, addUser.Email, addUser.Pass)
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
	}
	writer.Write(jsonData)
	//writer = utilities.JsonFetch(user, writer)
}
