package models

type AddTask struct {
	Task string `db:"task" json:"task"`
}
