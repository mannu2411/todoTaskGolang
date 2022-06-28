package helper

import (
	"github.com/todoTask/database"
	"github.com/todoTask/models"
)

func CreateTask(uid, task string) (string, error) {
	SQL := `INSERT INTO tasks(user_id,task) VALUES ($1, $2) RETURNING id;`
	var id string
	err := database.Tutorial.Get(&id, SQL, uid, task)
	if err != nil {
		return "", err
	}
	return id, err
}
func GetAllTasks(uid string) ([]models.AllTask, error) {
	// Not giving output when compared with sessionId
	SQL := `SELECT id, user_id, task, is_complete, created_at FROM tasks WHERE archived_at IS NULL;`
	var tasks []models.AllTask
	rows, errRow := database.Tutorial.Query(SQL)

	if errRow != nil {
		return tasks, errRow
	}
	//log.Printf(email + " herk")

	for rows.Next() {
		var u models.AllTask
		rows.Scan(&u.ID, &u.UserID, &u.Task, &u.IsComplete, &u.CreatedAt)

		if u.UserID == uid {
			tasks = append(tasks, u)
			//log.Printf(u.Email)
		}
	}
	return tasks, nil
}
func DeleteTask(tid string) (string, error) {
	//language=SQL
	SQL := `UPDATE tasks SET archived_at=CURRENT_TIMESTAMP WHERE id=$1 AND archived_at IS NULL RETURNING id;`

	var taskID string
	err := database.Tutorial.Get(&taskID, SQL, tid)
	if err != nil {
		return "", err
	}
	return taskID, nil
}
func CompleteTask(tid string) (string, error) {
	//language=SQL
	SQL := `UPDATE tasks SET is_complete=true WHERE id=$1 AND archived_at IS NULL RETURNING id;`

	var taskID string
	err := database.Tutorial.Get(&taskID, SQL, tid)
	if err != nil {
		return "", err
	}
	return taskID, nil
}
