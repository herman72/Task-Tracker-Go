package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)


type User struct {
	ID int
	email string
	password string
}

var userStorage []User

func main() {
	fmt.Println("Hello todo application")
	command := flag.String("command", "no command provided", "add, update, delete, mark-done, mark-in-progress")
	flag.Parse()

	for {
		runCommand(*command)
		fmt.Println("Enter command: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		scanner.Text()
	}

}

func runCommand(command string) {

	switch command {
	case "create-task":
		createTask()
	case "create-category":
		createCategory()
	case "register-user":
		registerUser()
	case "login":
		login()
	case "exit":
		os.Exit(0)
	default:
		fmt.Println("command is not valid", command)
	}

}

func createTask() {
	
	scanner := bufio.NewScanner(os.Stdin)

	var name, duedate, category string
		
	fmt.Println("pls enter the task title: ")
	scanner.Scan()
	name = scanner.Text()

	fmt.Println("pls enter the task duedate: ")
	scanner.Scan()
	duedate = scanner.Text()

	fmt.Println("pls enter the task category: ")
	scanner.Scan()
	category = scanner.Text()

	fmt.Println("Task title: , Task duedate: , Task category: ", name, duedate, category)

}

func registerUser(){
	scanner := bufio.NewScanner(os.Stdin)

	var id, email, password string
	fmt.Println("pls enter the User email: ")
	scanner.Scan()
	email = scanner.Text()

	fmt.Println("pls enter the User password: ")
	scanner.Scan()
	password = scanner.Text()

	id = email

	fmt.Println("User name:, User email:, User password:", id, email, password)

	user := User{
		ID: len(userStorage) + 1,
		email: email,
		password: password,
	}

	userStorage = append(userStorage, user)
}

func createCategory(){
	scanner := bufio.NewScanner(os.Stdin)

	var title, color string

	fmt.Println("pls enter the Category title: ")
	scanner.Scan()
	title = scanner.Text()

	fmt.Println("pls enter the Category color: ")
	scanner.Scan()
	color = scanner.Text()

	fmt.Println("Category title:, Category color:", title, color)
}

func login(){
	scanner := bufio.NewScanner(os.Stdin)

	var email, password string
	
	fmt.Println("pls enter the User email: ")
	scanner.Scan()
	email = scanner.Text()

	fmt.Println("pls enter the User password: ")
	scanner.Scan()
	password = scanner.Text()

	fmt.Println("User email:, User password", email, password)
}