package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
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

func updateTask(id int, description string) {
	tasks := loadTask()
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Description = description
			tasks[i].UpdatedAt = time.Now().Format(time.RFC3339)
			saveTasks(tasks)
			fmt.Printf("Task updated successfully (ID: %d)\n", id)

		}
	}

}

func deleteTask(id int) {
	tasks := loadTask()
	found := false
	for i := range tasks {
		if tasks[i].ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			saveTasks(tasks)
			fmt.Printf("Task deleted successfully (ID: %d)\n", id)
			found = true
			break

		}
	}
	if !found {
		fmt.Printf("Task not found (ID: %d)\n", id)
	}
}

func markTask(id int, status string) {
	tasks := loadTask()
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Status = status
			tasks[i].UpdatedAt = time.Now().Format(time.RFC3339)
			saveTasks(tasks)
			fmt.Printf("Task status with Id %d updated", id)
			break
		} else {
			fmt.Printf("Task with ID %d was not found", id)
		} 
	}
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
	case "update":
		if len(os.Args) < 3 {
			fmt.Println("Please provide a task description")
		} else {
			idStr := os.Args[2]
			id, _ := strconv.Atoi(idStr)
			description := os.Args[3]
			updateTask(id, description)
		}
	case "delete":
		if len(os.Args) < 2 {
			fmt.Printf("please provide the task ID")
		} else {
			idStr := os.Args[2]
			id, _ := strconv.Atoi(idStr)
			deleteTask(id)
		}
	case "mark-done":
		if len(os.Args) < 2 {
			fmt.Print("Please provide the task ID")
		} else {
			idStr := os.Args[2]		
			id, _ := strconv.Atoi(idStr)
			markTask(id, "done")
		}
	case "mark-in-progress":
		if len(os.Args) < 2 {
			fmt.Print("Please provide the task ID")
		} else {
			idStr := os.Args[2]		
			id, _ := strconv.Atoi(idStr)
			markTask(id, "in-progress")
		}

	default:
		fmt.Println("Unknown command:", command)
	}
	fmt.Println("Task Tracker CLI")

}
