package api

import (
	_ "github.com/lib/pq"

	//"database/sql"
	//"time"
)

type TaskObject struct {
	Owner int
	Text  string
	//DateCreated time.Date
	//DateDue     time.Date
}

func CreateNewTask(task *TaskObject) (err error) {
	return nil
}

func GetAllUserTasks(userId int) (err error) {
	return nil
}
