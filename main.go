package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

const taskFile = "task.json"

type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

func initTaskFile() {
	if _, err := os.Stat(taskFile); os.IsNotExist(err) {
		emptyTask := []Task{}
		file, _ := json.MarshalIndent(emptyTask, "", " ")
		_ = os.WriteFile(taskFile, file, 0644)
	}
}

func loadTask() []Task {
	file, _ := os.ReadFile(taskFile)
	var task []Task
	_ = json.Unmarshal(file, &task)
	return task
}

func saveTasks(tasks []Task) {
	file, _ := json.MarshalIndent(tasks, "", " ")
	_ = os.WriteFile(taskFile, file, 0644)
}

func addTask(description string) {
	tasks := loadTask()
	newTask := Task{
		ID:          len(tasks) + 1,
		Description: description,
		Status:      "todo",
		CreatedAt:   time.Now().Format(time.RFC3339),
		UpdatedAt:   time.Now().Format(time.RFC3339),
	}
	tasks = append(tasks, newTask)
	saveTasks(tasks)
	fmt.Printf("Task added successfully (ID: %d)\n", newTask.ID)
}

func main() {
	initTaskFile()

	if len(os.Args) < 2 {
		fmt.Println("Expected 'add' command with task description")
		return
	}
	
	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
            fmt.Println("Please provide a task description")
        } else {
            addTask(os.Args[2])
        }
	default:
		fmt.Println("Unknown command:", command)
	}
	fmt.Println("Task Tracker CLI")

}
