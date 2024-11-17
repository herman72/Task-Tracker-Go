package main

import (
	"os"
	"testing"
)

func TestAddTask(t *testing.T){
	err := os.WriteFile(taskFile, []byte("[]"), 0644)
	if err != nil {
		t.Fatalf("Failed to setup task: %v", err)
	}

	addTask("Test task")

	tasks := loadTask()

	if len(tasks) != 1 {
		t.Fatalf("Expected 1 task, but got %d", len(tasks))
	}

	if tasks[0].Description != "Test task" {
		t.Errorf("Expected task description to be 'Test task', but got '%s'", tasks[0].Description)
	}

	if tasks[0].Status != "todo" {
		t.Errorf("Expected task status to be 'todo', but got '%s'", tasks[0].Status)
	}
}

func TestInitTaskFile(t *testing.T){
	err := os.Remove(taskFile)
	if err != nil {
		t.Fatalf("Failed to remove task file: %v", err)
	}

	initTaskFile()

	_, err = os.Stat(taskFile)
	if os.IsNotExist(err) {
		t.Fatalf("Task file was not created")
	}

	initTaskFile()

	if _, err := os.Stat(taskFile); os.IsNotExist(err) {
		t.Fatalf("Expected task file to be created, but it does not exist")
	}

	tasks  := loadTask()

	if len(tasks) != 0 {
		t.Fatalf("Expected no task, but got %d", len(tasks))
	}

	
}