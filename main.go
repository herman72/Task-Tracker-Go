package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"task-tracker-cli/contract"
	"task-tracker-cli/entity"
	"task-tracker-cli/filestore"
	"task-tracker-cli/constant"
	"crypto/md5"
	"encoding/hex"
)


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
	userStorage     []entity.User
	taskStorage     []Task
	categoryStorage []Category

	authenticatedUser *entity.User
	serialiazatinMode string
)

const (
	userStoragePath        = "user.txt"
	
)

func main() {
	fmt.Println("Hello todo application")
	serilizedMode := flag.String("serilize-mode", constant.OldOneSerilizationMode, "serilization mode for writing data")
	command := flag.String("command", "no command provided", "add, update, delete, mark-done, mark-in-progress")
	flag.Parse()

	// loadUserStorageFromFile(*serilizedMode)

	

	switch *serilizedMode {
	case constant.OldOneSerilizationMode:
		serialiazatinMode = constant.OldOneSerilizationMode
	default:
		serialiazatinMode = constant.JsonSerializationMode
	}
	var userFileStore = filestore.New(userStoragePath, serialiazatinMode)

	
	users := userFileStore.Load()
	userStorage = append(userStorage, users...)

	for {
		runCommand(userFileStore, *command)
		fmt.Println("Enter command: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		*command = scanner.Text()
	}

}

func runCommand(store contract.UserWriteStore, command string) {

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
		registerUser(store)
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



func registerUser(store contract.UserWriteStore) {
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
	
	user := entity.User{
		ID:       len(userStorage) + 1,
		Name:     name,
		Email:    email,
		Password: hashThePassword(password),
	}

	userStorage = append(userStorage, user)
	// writeUserToFile(user)
	store.Save(user)
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

// func loadUserStorageFromFile(serialiazatinMode string) {
	
// }



func hashThePassword(password string) string {

	hash := md5.Sum([]byte(password))

	return hex.EncodeToString(hash[:])

}

