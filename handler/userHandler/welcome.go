package userHandler

import (
	"encoding/json"
	"net/http"

	"github.com/todoTask/database/helper"
	"github.com/todoTask/models"
)

func isErr(err error, typeErr string) bool {
	return err.Error() == "pq: duplicate key value violates unique constraint "+typeErr
}

func Test(writer http.ResponseWriter, request *http.Request) {
	_, err := helper.CreateUser("test1", "test1@test.com", "test1")
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	var user = models.AddUser{"test1", "test1@test.com", "test1"}
	jsonData, jsonErr := json.Marshal(user)
	if jsonErr != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.Write(jsonData)
}
