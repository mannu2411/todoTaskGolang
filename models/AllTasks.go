package models

import (
	"time"
)

type AllTask struct {
	ID         string    `db:"id" json:"id"`
	Email      string    `db:"email" json:"email"`
	Task       string    `db:"task" json:"task"`
	IsComplete bool      `db:"iscomplete" json:"iscomplete"`
	CreatedAt  time.Time `db:"created_at" json:"createdAt"`
}
