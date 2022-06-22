package models

type Credentials struct {
	Email string `db:"email" json:"email"`
	Pass  string `db:"pass" json:"pass"`
}
