package main

import (
	"fmt"
)

// Type to be returned when there is an error parsing a form
type FormValidationError struct {
	What string
}

// Error string function for FormValidationError
func (err FormValidationError) Error() string {
	return fmt.Sprintf("FormValidationError:: %v", err.What)
}

// Struct used for registering a new user
type CreateAccountForm struct {
	Email     string
	FirstName string
	LastName  string
	Password  string
	Password2 string
}

// Validate function for CreateAccountForm
func (form *CreateAccountForm) Validate() (err error) {
	// select query to check if email is unique
	if form.Email == "" {
		return FormValidationError{
			"Email cannot be blank!",
		}
	} else if form.FirstName == "" {
		return FormValidationError{
			"First name cannot be blank!",
		}
	} else if form.LastName == "" {
		return FormValidationError{
			"Last name cannot be blank",
		}
	} else if form.Password == "" || form.Password != form.Password2 {
		return FormValidationError{
			"Passwords do not match!",
		}
	} else {
		return nil
	}
}

// Struct used for logging in a user
type LoginForm struct {
	Email      string
	Password   string
	RememberMe bool
}

// Validate function for LoginForm
func (form *LoginForm) Validate() (err error) {
	// select query to check email
	if form.Email == "" {
		return FormValidationError{
			"Email cannot be blank!",
		}
	} else if form.Password == "" {
		return FormValidationError{
			"Passwords cannot be blank!",
		}
	} else {
		return nil
	}
}
