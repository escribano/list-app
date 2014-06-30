package api

import (
	_ "github.com/lib/pq"

	//"database/sql"
	"fmt"
)

func CreateNewTag(tag string) (err error) {
	stmt, err := DB.Prepare("INSERT INTO tags (tag_text) VALUES ($1);")
	if err != nil {
		fmt.Println("ERROR preparing statement: ", err)
		return err
	}

	results, err := stmt.Exec(tag)
	if err != nil {
		fmt.Println("ERROR inserting new tag: ", err)
		return err
	}
	fmt.Println(results)
	return nil
}

func CreateTagMap(taskId int, tagText string) (err error) {
	// get the tagId
	stmt, err := DB.Prepare("select tag_id from tags where tag_text = $1;")
	if err != nil {
		fmt.Println("ERROR preparing statement select: ", err)
	}

	row := stmt.QueryRow(tagText)
	var tagId int
	if err = row.Scan(&tagId); err != nil {
		fmt.Println("ERROR retrieving tag_id: ", err)
	}
	//row.Close()

	stmt, err = DB.Prepare("INSERT INTO tagmap (task_id, tag_id) VALUES ($1, $2);")
	if err != nil {
		fmt.Println("ERROR preparing statement insert: ", err)
		return err
	}

	results, err := stmt.Exec(taskId, tagId)
	if err != nil {
		fmt.Println("ERROR inserting new tag: ", err)
		return err
	}
	fmt.Println(results)
	return nil
}
