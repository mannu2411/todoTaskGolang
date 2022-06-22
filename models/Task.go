package models

type Task struct {
	ID        string `db:"id" json:"id"`
	SessionID string `db:"sessionid" json:"sessionid"`
}
