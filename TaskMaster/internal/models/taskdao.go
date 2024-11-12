package models

import (
	"TaskMaster/internal/db"
)

type Task struct {
	ID     int    `json:"id"`
	UserID int    `json:"userId"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

// CreateTask inserts a new task into the database
func CreateTask(userID int, title string, status string) error {
	var databaseTask = db.GetDB()
	_, err := databaseTask.Exec("INSERT INTO tasks (user_id, title, status) VALUES ($1, $2, $3)", userID, title, status)
	return err
}

// GetTasks retrieves all tasks from the database
func GetTasks() ([]Task, error) {
	var databaseTask = db.GetDB()
	rows, err := databaseTask.Query("SELECT id, user_id, title, status FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		if err := rows.Scan(&task.ID, &task.UserID, &task.Title, &task.Status); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}
