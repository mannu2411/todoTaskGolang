package userHandler

import (
	"encoding/json"
	"github.com/todoTask/middlewareUtils"
	"log"
	"net/http"

	"github.com/todoTask/database/helper"
)

func DeleteUser(writer http.ResponseWriter, request *http.Request) {

	/*	var deleteUser models.UpdateUser
			addErr := json.NewDecoder(request.Body).Decode(&deleteUser)
			if addErr != nil {
				writer.WriteHeader(http.StatusBadRequest)
				return
			}
		session := request.Header.Values("session_token")
		sessionId := session[0]
	*/
	userID := middlewareUtils.UserFromContext(request.Context()).ID
	err := helper.DeleteSession(userID)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	user := middlewareUtils.UserFromContext(request.Context()).ID
	userId, err := helper.DeleteUser(user)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Printf("User: " + userId + "has been deleted")
	jsonData, jsonErr := json.Marshal(user)
	if jsonErr != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.Write(jsonData)
}
