package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"crypto/md5"
	"encoding/hex"
)

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}

type Task struct {
	ID         int
	Title      string
	Duedate    string
	CategoryID int
	IsDone     bool
	UserId     int
}

type Category struct {
	ID     int
	Title  string
	Color  string
	UserId int
}

var (
	userStorage     []User
	taskStorage     []Task
	categoryStorage []Category

	authenticatedUser *User
	serialiazatinMode string
)

const (
	userStoragePath        = "user.txt"
	oldOneSerilizationMode = "oldone"
	JsonSerializationMode  = "json"
)

func main() {
	fmt.Println("Hello todo application")
	serilizedMode := flag.String("serilize-mode", oldOneSerilizationMode, "serilization mode for writing data")
	command := flag.String("command", "no command provided", "add, update, delete, mark-done, mark-in-progress")
	flag.Parse()

	loadUserStorageFromFile(*serilizedMode)

	switch *serilizedMode {
	case oldOneSerilizationMode:
		serialiazatinMode = oldOneSerilizationMode
	default:
		serialiazatinMode = JsonSerializationMode
	}

	for {
		runCommand(*command)
		fmt.Println("Enter command: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		*command = scanner.Text()
	}

}

func runCommand(command string) {

	if command != "register-user" && command != "exit" && authenticatedUser == nil {
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

	categoryID, err := strconv.Atoi(category)

	if err != nil {
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
		ID:         len(taskStorage) + 1,
		Title:      title,
		Duedate:    duedate,
		CategoryID: categoryID,
		IsDone:     false,
		UserId:     authenticatedUser.ID,
	}

	taskStorage = append(taskStorage, task)

	fmt.Println("Task title: , Task duedate: , Task category: ", title, duedate, category)

}

func registerUser() {
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
		ID:       len(userStorage) + 1,
		Name:     name,
		Email:    email,
		Password: hashThePassword(password),
	}

	userStorage = append(userStorage, user)
	writeUserToFile(user)

}

func createCategory() {
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
		ID:     len(categoryStorage) + 1,
		Title:  title,
		Color:  color,
		UserId: authenticatedUser.ID,
	}

	categoryStorage = append(categoryStorage, category)

}

func login() {
	scanner := bufio.NewScanner(os.Stdin)

	var email, password string

	fmt.Println("pls enter the User email: ")
	scanner.Scan()
	email = scanner.Text()

	fmt.Println("pls enter the User password: ")
	scanner.Scan()
	password = scanner.Text()

	for _, user := range userStorage {
		if user.Email == email && user.Password == hashThePassword(password) {
			authenticatedUser = &user
			fmt.Println("you are logged in")

			break
		}

	}
	if authenticatedUser == nil {
		fmt.Println("the email or password is not correct")
	}

}

func listTask() {
	for _, task := range taskStorage {
		if task.ID == authenticatedUser.ID {
			fmt.Println(task)
		}

	}
}

func loadUserStorageFromFile(serialiazatinMode string) {
	file, err := os.Open(userStoragePath)
	
	if err != nil {
		fmt.Println("there is no file", err)
	}

	var data = make([]byte, 1024)
	_, oErr := file.Read(data)

	if oErr != nil {
		fmt.Println("can't read from ", oErr)
	}

	var dataString = string(data)
	dataString = strings.Trim(dataString, "\n")
	userSlice := strings.Split(dataString, "\n")

	for _, u := range userSlice {
		
		var userStruct = User{}

		switch serialiazatinMode {
		case oldOneSerilizationMode:
			var dErr error
			userStruct, dErr = deSerilizedOldOne(u)

			if dErr != nil {
				fmt.Println("cant desrilized user record to user struct", dErr)
				return
			}

		case JsonSerializationMode:
			if u[0] != '{' && u[len(u)-1] != '}' {
				continue
			}

			uErr := json.Unmarshal([]byte(u), &userStruct)
			if uErr != nil {
				fmt.Println("cant desrilized user record to user struct from json mode", uErr)
				return
			}
		}
		userStorage = append(userStorage, userStruct)

	}
}

func writeUserToFile(user User) {
	var file *os.File

	file, err := os.OpenFile(userStoragePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println("can't write file", err)
	}
	defer file.Close()
	// Serialized the user struct
	var data []byte
	if serialiazatinMode == oldOneSerilizationMode {
		data = []byte(fmt.Sprintf("id: %d, name: %s, email: %s, password: %s\n",
			user.ID, user.Name, user.Email, user.Password))

	} else if serialiazatinMode == JsonSerializationMode {
		var jErr error
		data, jErr = json.Marshal(user)
		data = append(data, []byte("\n")...)
		if err != nil {
			fmt.Println("can't mashal user struct to json", jErr)
			return
		}
	} else {
		fmt.Println("invalid serilization mode")

		return
	}

	_, wErr := file.Write([]byte(data))
	if wErr != nil {
		fmt.Println("can't write file", wErr)
	}

}

func deSerilizedOldOne(userStr string) (User, error) {
	if userStr == "" {
		return User{}, errors.New("use string is empty ")
	}
	userFields := strings.Split(userStr, ",")
	var user = User{}
	for _, field := range userFields {
		values := strings.Split(field, ": ")
		if len(values) != 2 {
			fmt.Printf("invalid field format: %v\n", field)
			continue
		}
		fieldName := strings.ReplaceAll(values[0], " ", "")
		fieldValue := values[1]

		switch fieldName {
		case "id":
			id, err := strconv.Atoi(fieldValue)
			if err != nil {
				fmt.Println("error in ")

				return User{}, errors.New("strconv error")
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
	return user, nil

}

func hashThePassword(password string) string {

	hash := md5.Sum([]byte(password))

	return hex.EncodeToString(hash[:])

}