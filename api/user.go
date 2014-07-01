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

// api function for creating a new user
// given the user's imformation create and add a user to the database.
// if there is an error return the error.
func CreateNewUser(email, first, last, hash, salt string) (err error) {
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

// api function for getting a user from an email
// given a user's email gets the user's information and returns it.
// if there is an error return the error as well.
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

// api function for updating a user
func UpdateUser(userId int, fields map[string]string) (err error) {
	return nil
}

// api function for deleting a user
func DeleteUser(userId int) (err error) {
	return nil
}
