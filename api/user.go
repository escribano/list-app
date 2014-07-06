package api

import (
	"fmt"
)

type UserObject struct {
	UserId    int
	FirstName string
	LastName  string
	Email     string
	Salt      string
	Hash      string
}

// CreateNewUser is an api function for creating a new user.
// Given the user's form information, add a user to the database.
// Tjis functions does not do any validation, that should all be done by the handler.
func CreateNewUser(email, first, last, hash, salt string) error {
	stmt, err := DB.Prepare("INSERT INTO users (email, first_name, last_name, hash, salt) VALUES ($1, $2, $3, $4, $5);")
	if err != nil {
		fmt.Println("ERROR preparing statement: ", err)
		return err
	}
	results, err := stmt.Exec(email, first, last, hash, salt)
	if err != nil {
		fmt.Println("ERROR inserting new user: ", err)
		return err
	}
	fmt.Println(results)
	return nil
}

// GetUser is a general api function for getting a user from an email.
// Given a user's email gets the user's information and returns it as a UserObject.
func GetUser(email string) (UserObject, error) {
	stmt, err := DB.Prepare("SELECT user_id, first_name, last_name, email, hash, salt FROM users WHERE email = $1;")
	if err != nil {
		fmt.Println("ERROR preparing statement: ", err)
		return UserObject{}, err
	}

	var user UserObject
	err = stmt.QueryRow(email).Scan(&user.UserId, &user.FirstName,
		&user.LastName, &user.Email, &user.Hash, &user.Salt)
	if err != nil {
		fmt.Println("ERROR selecting all users: ", err)
		return user, err
	}
	return user, nil
}

// UpdateUser is an api function for updating parts or all of a user's information.
func UpdateUser(userId int, fields map[string]string) error {
	return nil
}

// DeleteUser is an api function for deleting a user and all associated tasks/data.
func DeleteUser(userId int) error {
	return nil
}
