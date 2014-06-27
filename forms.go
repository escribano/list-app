package main

type CreateAccountForm struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
	Password2 string
}

type LoginForm struct {
	Email      string
	Password   string
	RememberMe bool
}
