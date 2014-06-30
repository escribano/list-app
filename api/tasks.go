package api

import (
	_ "github.com/lib/pq"

	"database/sql"
	"fmt"
	"time"
)

type TaskObject struct {
	TaskId      int
	Owner       int
	Text        string
	DateCreated time.Time
	//DateDue     time.Date
}

// Create a new task
// TODO: Make this (and the database) accept a due date and possibly reminders
func CreateTask(text string, owner int) (err error) {
	stmt, err := DB.Prepare("INSERT INTO tasks (task_text, task_owner, date_created) VALUES ($1, $2, $3);")
	if err != nil {
		fmt.Println("ERROR preparing statement: ", err)
		return err
	}

	results, err := stmt.Exec(text, owner, time.Now())
	if err != nil {
		fmt.Println("ERROR inserting new task: ", err)
		return err
	}
	return nil
}

// Delete a given task
func DeleteTask(taskId int) (err error) {
	stmt, err := DB.Prepare("DELETE FROM tasks WHERE task_id = $1;")
	if err != nil {
		fmt.Println("ERROR preparing statement: ", err)
		return err
	}

	results, err := stmt.Exec(taskId)
	if err != nil {
		fmt.Println("ERROR deleting task: ", err)
		return err
	}
	return nil
}

// Takes a TaskObject. This call will replace task_text and due_date with that in the object
// TODO: Optimize and make not stupid
func UpdateTask(updatedTask TaskObject) (err error) {
	stmt, err := DB.Prepare("UPDATE tasks SET task_text = $1 WHERE task_id = $2;")
	if err != nil {
		fmt.Println("ERROR preparing statement: ", err)
		return err
	}

	results, err := stmt.Exec(updatedTask.Text, updatedTask.TaskId)
	if err != nil {
		fmt.Println("ERROR updating task: ", err)
		return err
	}
	return nil
}

func GetTaskById(taskId int) (err error) {
	return nil
}

// Get all a user's tasks. Default api call for a freshly logged in user
func GetAllUserTasks(userId int) ([]TaskObject, error) {
	var tasks []TaskObject
	stmt, err := DB.Prepare("SELECT * FROM tasks WHERE task_owner = $1;")
	if err != nil {
		fmt.Println("ERROR preparing statement: ", err)
		return tasks, err
	}

	rows, err := stmt.Query(userId)
	if err != nil {
		fmt.Println("ERROR selecting all tasks: ", err)
		return tasks, err
	}
	defer rows.Close()

	for rows.Next() {
		var task TaskObject
		if err := rows.Scan(&task.TaskId, &task.Owner, &task.Text, &task.DateCreated); err != nil {
			fmt.Println("ERROR scanning tasks: ", err)
		}
		tasks = append(tasks, task)
	}
	if err := rows.Err(); err != nil {
		fmt.Println("Row errorr?", err)
	}
	return tasks, nil
}

// Verify that taskId is owned by taskOwner
// This is a general purpose api call that should be used before any alteration of an existing task
func VerifyTaskOwner(taskId, taskOwner int) bool {
	stmt, err := DB.Prepare("Select task_id, task_owner From tasks where task_id = $1 AND task_owner = $2;")
	if err != nil {
		fmt.Println("ERROR preparing statement: ", err)
	}
	var dbId, dbOwner int
	row := stmt.QueryRow(taskId, taskOwner)
	if err = row.Scan(&dbId, &dbOwner); err == sql.ErrNoRows {
		fmt.Println("Invalid task access, user does not own task")
		return false
	} else if err != nil {
		fmt.Println("ERROR: ", err)
		return false
	}
	return true
}
