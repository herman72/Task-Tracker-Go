package main

import (
	"encoding/json"
	"fmt"
	"os"
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

func main() {
	initTaskFile()
	fmt.Println("Task Tracker CLI")

}
