package helper

import (
	"database/sql"

	"github.com/todoTask/database"
	"github.com/todoTask/models"
)

func CreateUser(name, email, pass string) (string, error) {
	// language=SQL
	SQL := `INSERT INTO users(name, email, pass) VALUES ($1, $2, $3) RETURNING id;`
	var userID string
	err := database.Tutorial.Get(&userID, SQL, name, email, pass)
	if err != nil {
		return "", err
	}
	return userID, nil
}

func GetUser(userID string) (*models.User, error) {
	// language=SQL
	SQL := `SELECT id, name, email, created_at, archived_at FROM users WHERE archived_at IS NULL AND id = $1`
	var user models.User
	err := database.Tutorial.Get(&user, SQL, userID)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &user, nil
}

func GetAllUser() ([]models.User, error) {
	// language=SQL
	SQL := `SELECT id, name, email, created_at, archived_at FROM users WHERE archived_at IS NULL;`
	var users []models.User
	rows, errRow := database.Tutorial.Queryx(SQL)
	if errRow != nil {
		return users, errRow
	}
	for rows.Next() {
		var u models.User
		rows.Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt, &u.ArchivedAt)
		users = append(users, u)
	}
	return users, nil
}

func UpdateUser(name, email, id string) (string, error) {
	//language=SQL
	SQL := `UPDATE users SET name=$1, email=$2 WHERE id=$3 RETURNING id;`
	var userID string
	err := database.Tutorial.Get(&userID, SQL, name, email, id)
	if err != nil {
		return "", err
	}
	return userID, nil
}

func DeleteUser(uid string) (string, error) {
	//language=SQL
	SQL := `UPDATE users SET archived_at=CURRENT_TIMESTAMP WHERE id=$1 AND archived_at IS NULL RETURNING id;`

	var userID string
	err := database.Tutorial.Get(&userID, SQL, uid)
	if err != nil {
		return "", err
	}
	return userID, nil
}
