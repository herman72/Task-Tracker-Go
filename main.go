package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	CategoryID int
	IsDone bool
	UserId int
}

type Category struct{
	 ID int
	 Title string
	 Color string
	 UserId int
}

var userStorage []User
var authenticatedUser *User

var taskStorage []Task
var categoryStorage []Category

const userStoragePath = "user.txt"

func main() {

	loadUserStorageFromFile()

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

	fmt.Println("pls enter the task category id: ")
	scanner.Scan()
	category = scanner.Text()

	categoryID, err :=strconv.Atoi(category)

	if err !=nil{
		fmt.Printf("category id is not valid %v\n", err)
		
		return
	}
	isFound := false
	for _, c := range categoryStorage {
		if c.ID == categoryID && c.UserId == authenticatedUser.ID {
			isFound = true
			break
		}
	}
	if !isFound {
		fmt.Printf("category id is not valid, \n")
	}

	task := Task{
		ID: len(taskStorage) + 1,
		Title: title,
		Duedate: duedate,
		CategoryID: categoryID ,
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

	var file *os.File

	file, err := os.OpenFile(userStoragePath, os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println("can't write file", err)
	}

	data := fmt.Sprintf("id: %d, name: %s, email: %s, password: %s\n", 
	user.ID, user.Name, user.Email, user.Password)

	file.Write([]byte(data))
	file.Close()
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

	category := Category{
		ID: len(categoryStorage) + 1,
		Title: title,
		Color: color,
		UserId: authenticatedUser.ID,
	}

	categoryStorage = append(categoryStorage, category)

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

func loadUserStorageFromFile(){
	file, err := os.Open(userStoragePath)

	if err != nil {
		fmt.Println("there is no file", err)
	}

	var data = make([]byte, 1024)
	_, oErr :=file.Read(data)

	if oErr != nil {
		fmt.Println("can't read from ", oErr)
	}

	var dataString = string(data)
	dataString = strings.Trim(dataString, "\n")
	userSlice := strings.Split(dataString, "\n")
	for _, u := range userSlice {
		userFields := strings.Split(u, ",")
		var user = User{}
		for _, field := range userFields {
			values := strings.Split(field, ": ")
			fieldName := strings.ReplaceAll(values[0], " ", "")
			fieldValue := values[1]

			

			switch fieldName {
			case "id":
				id, err := strconv.Atoi(fieldValue)
				if err != nil {
					fmt.Println("error in ")
					
					return
				}
				user.ID = id
			
			case "name":
				user.Name = fieldValue
			case "email":
				user.Email = fieldValue
			case "password":
				user.Password = fieldValue
			}



		}
		fmt.Printf("user %v\n", user)

	}
}