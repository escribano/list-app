package api

import (
	_ "github.com/lib/pq"

	//"database/sql"
	"fmt"
)

func CreateNewTag(tag string) error {
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
