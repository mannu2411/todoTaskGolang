package helper

import (
	"github.com/todoTask/database"
	"github.com/todoTask/models"
)

func CreateTask(sessionId, task string) (string, error) {
	SQL := `INSERT INTO tasks(sessionid,task) VALUES ($1, $2) RETURNING id;`
	var id string
	err := database.Tutorial.Get(&id, SQL, sessionId, task)
	if err != nil {
		return "", err
	}
	return id, err
}
func GetAllTasks(sessionId string) ([]models.AllTask, error) {
	// Not giving output when compared with sessionId
	SQL := `SELECT sessionid, id, task, iscomplete, created_at FROM tasks WHERE archived_at IS NULL;`
	var tasks []models.AllTask
	rows, errRow := database.Tutorial.Queryx(SQL)

	if errRow != nil {
		return tasks, errRow
	}
	for rows.Next() {
		var u models.AllTask
		rows.Scan(&u.SessionID, &u.ID, &u.Task, &u.IsComplete, &u.CreatedAt)
		if u.SessionID == sessionId {
			tasks = append(tasks, u)
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
	SQL := `UPDATE tasks SET iscomplete=true WHERE id=$1 AND archived_at IS NULL RETURNING id;`

	var taskID string
	err := database.Tutorial.Get(&taskID, SQL, tid)
	if err != nil {
		return "", err
	}
	return taskID, nil
}
