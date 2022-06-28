package models

type UpdateUser struct {
	Name  string `db:"name" json:"name"`
	Email string `db:"email" json:"email"`
}
