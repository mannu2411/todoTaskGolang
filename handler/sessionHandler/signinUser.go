package sessionHandler

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/todoTask/database/helper"
	"github.com/todoTask/models"
)

func SignInUser(writer http.ResponseWriter, request *http.Request) {
	var creds models.Credentials
	err := json.NewDecoder(request.Body).Decode(&creds)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	expectPass, err := helper.GetPass(creds.Email)

	if err != nil || expectPass != creds.Pass {
		writer.WriteHeader(http.StatusUnauthorized)
	}

	expireAt := time.Now().Add(360 * time.Second)
	//uid, err := uuid.NewV4()
	if err != nil {
		log.Fatalf("failed to generate UUID: %v", err)
	}
	//id := uid.String()
	sessionId, err := helper.GetSession(creds.Email)
	if sessionId == "" {
		sessionId, err = helper.CreateSession(creds.Email, expireAt)
	}

	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	http.SetCookie(writer, &http.Cookie{
		Name:    "session_token",
		Value:   sessionId,
		Expires: expireAt,
	})

}
