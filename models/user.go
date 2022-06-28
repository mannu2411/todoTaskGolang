package models

import (
	"time"

	"github.com/volatiletech/null"
)

type User struct {
	ID         string    `db:"id" json:"id"`
	Name       string    `db:"name" json:"name"`
	Email      string    `db:"email" json:"email"`
	CreatedAt  time.Time `db:"created_at" json:"createdAt"`
	ArchivedAt null.Time `db:"archived_at" json:"archivedAt"`
}

const ActiveUser string = "active_user"

type Authentication struct {
	Email string `json:"email"`
	Pass  string `json:"pass"`
}

type Token struct {
	Email       string `json:"email"`
	TokenString string `json:"token"`
}
