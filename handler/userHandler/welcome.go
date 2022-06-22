package userHandler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/todoTask/database/helper"
)

func isErr(err error, typeErr string) bool {
	return err.Error() == "pq: duplicate key value violates unique constraint "+typeErr
}

func Greet(writer http.ResponseWriter, request *http.Request) {
	userID, err := helper.CreateUser("test1", "test1@test.com", "test1")
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
