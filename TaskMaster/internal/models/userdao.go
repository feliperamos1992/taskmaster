package models

import (
	"TaskMaster/internal/db"
	"database/sql"
	"errors"
	"fmt"
)

type User struct {
	ID   int
	Name string
}

type UserTask struct {
	ID     int    `json:"id"`
	UserID int    `json:"userId"`
	Name   string `json:"name"`
	Tasks  []Task `json:"tasks"`
}

// CreateUser inserts a new user into the database
func CreateUser(name string) error {
	var databaseUser = db.GetDB()
	_, err := databaseUser.Exec("INSERT INTO users (name) VALUES ($1)", name)
	return err
}

func GetUsers(id int) (*User, error) {
	var databaseUser = db.GetDB()
	row := databaseUser.QueryRow("SELECT id, name FROM users WHERE id = $1", id)

	var user User
	if err := row.Scan(&user.ID, &user.Name); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("no user with id: %d", id)
		}
		return nil, err
	}
	return &user, nil
}

func GetUsersAndTasks(userId int) (*UserTask, error) {
	var databaseUser = db.GetDB()
	rows, err := databaseUser.Query(`SELECT u.id, u.name, t.id, t.title FROM users u LEFT JOIN tasks t ON u.id = t.user_id WHERE u.id = $1`, userId)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var userTask UserTask
	userTask.Tasks = []Task{}

	for rows.Next() {
		var task Task
		var taskID sql.NullInt32
		var taskTitle sql.NullString

		if err := rows.Scan(&userTask.ID, &userTask.Name, &taskID, &taskTitle); err != nil {
			return nil, err
		}

		userTask.UserID = userId

		// Verifique se os valores são válidos antes de atribuí-los
		if taskID.Valid {
			task.ID = int(taskID.Int32)
		}
		if taskTitle.Valid {
			task.Title = taskTitle.String
		}

		if taskID.Valid { // Apenas adicione a task se taskID não for NULL
			userTask.Tasks = append(userTask.Tasks, task)
		}
	}

	if len(userTask.Tasks) == 0 && userTask.ID == 0 {
		return nil, fmt.Errorf("no user with id: %d", userId)
	}

	return &userTask, nil
}
