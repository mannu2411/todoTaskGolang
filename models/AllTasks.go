package models

import (
	"time"
)

type AllTask struct {
	ID         string    `db:"id" json:"id"`
	UserID     string    `db:"user_id" json:"user_id"`
	Task       string    `db:"task" json:"task"`
	IsComplete bool      `db:"is_complete" json:"is_complete"`
	CreatedAt  time.Time `db:"created_at" json:"createdAt"`
}
