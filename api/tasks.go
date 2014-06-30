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
	fmt.Println(results)
	return nil
}

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

// verify that taskId is owned by taskOwner
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
