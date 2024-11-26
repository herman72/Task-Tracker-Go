package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)


type User struct {
	ID int
	Name string
	Email string
	Password string
}

type Task struct {
	ID int
	Title string
	Duedate string
	Category string
	IsDone bool
	UserId int
}

var userStorage []User
var taskStorage []Task
var authenticatedUser *User

func main() {

	fmt.Println("Hello todo application")
	command := flag.String("command", "no command provided", "add, update, delete, mark-done, mark-in-progress")
	flag.Parse()
		
	for {
		runCommand(*command)
		fmt.Println("Enter command: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		*command = scanner.Text()
	}

}

func runCommand(command string) {

	if command != "register-user" && command != "exit" && authenticatedUser == nil{ 
		login()

		if authenticatedUser == nil {
			return
		}
	}

	switch command {
	case "create-task":
		createTask()
	case "list-task":
		listTask()
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

	var title, duedate, category string
		
	fmt.Println("pls enter the task title: ")
	scanner.Scan()
	title = scanner.Text()

	fmt.Println("pls enter the task duedate: ")
	scanner.Scan()
	duedate = scanner.Text()

	fmt.Println("pls enter the task category: ")
	scanner.Scan()
	category = scanner.Text()

	task := Task{
		ID: len(taskStorage) + 1,
		Title: title,
		Duedate: duedate,
		Category: category ,
		IsDone: false,
		UserId: authenticatedUser.ID,
		

	}

	taskStorage = append(taskStorage, task)

	fmt.Println("Task title: , Task duedate: , Task category: ", title, duedate, category)

}

func registerUser(){
	scanner := bufio.NewScanner(os.Stdin)

	var id, email, name, password string
	
	fmt.Println("pls enter the User name: ")
	scanner.Scan()
	name = scanner.Text()
	
	fmt.Println("pls enter the User email: ")
	scanner.Scan()
	email = scanner.Text()

	fmt.Println("pls enter the User password: ")
	scanner.Scan()
	password = scanner.Text()

	id = email

	fmt.Println("User name:, User email:, User password:", id, name, email, password)

	user := User{
		ID: len(userStorage) + 1,
		Name: name,
		Email: email,
		Password: password,
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

		for _, user := range userStorage {
			if user.Email == email && user.Password == password {
				authenticatedUser = &user
				fmt.Println("you are logged in")

				break
			} 
			
		}

		if authenticatedUser == nil {
			fmt.Println("User not found")
			
		}


	fmt.Println("User email:, User password", email, password)
}

func listTask(){
	for _,task := range taskStorage{
		if task.ID == authenticatedUser.ID {
			fmt.Println(task)
		}
		
	}
}