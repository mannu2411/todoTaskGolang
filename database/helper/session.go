package helper

import (
	"time"

	"github.com/todoTask/database"
)

func GetPass(email string) (string, error) {
	SQL := `SELECT pass FROM users WHERE email=$1;`
	var pass string
	err := database.Tutorial.Get(&pass, SQL, email)
	if err != nil {
		return "", err
	}
	return pass, nil
}

func CreateSession(email string, end_at time.Time) (string, error) {

	SQL := `INSERT INTO session( email, end_at) VALUES ($1, $2) RETURNING id;`
	var sessionID string
	err := database.Tutorial.Get(&sessionID, SQL, email, end_at)
	if err != nil {
		return "", err
	}
	return sessionID, nil
}

/*
func GetSession(email string) (string, error) {
	SQL := `SELECT id FROM session WHERE email=$1;`
	var sid string
	err := database.Tutorial.Get(&sid, SQL, email)
	if err != nil {
		return "", err
	}
	return sid, nil
}

func UpdateSession(end_at time.Time, id string) error {
	//language=SQL
	SQL := `UPDATE session SET end_at=$1 WHERE id=$2 RETURNING id;`
	var userID string
	err := database.Tutorial.Get(&userID, SQL, end_at, id)
	if err != nil {
		return err
	}
	return nil
}
*/
func IsExpired(id string) (bool, error) {
	SQL := `SELECT end_at FROM session WHERE id=$1;`
	var expireAt time.Time
	err := database.Tutorial.Get(&expireAt, SQL, id)
	if err != nil {
		return true, err
	}
	if !expireAt.After(time.Now()) {
		return true, nil
	}
	return false, nil
}

func SessionExist(sessionId string) (bool, error) {
	SQL := `SELECT email FROM session WHERE id=$1;`
	var email string
	err := database.Tutorial.Get(&email, SQL, sessionId)
	if err != nil {
		return false, err
	}
	return true, nil
}

func DeleteSession(uid string) error {
	//language=SQL
	SQL := `UPDATE session SET end_at=CURRENT_TIMESTAMP WHERE id=$1 RETURNING id;`

	var taskID string
	err := database.Tutorial.Get(&taskID, SQL, uid)
	if err != nil {
		return err
	}
	return nil
}
