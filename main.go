package main

import (
	"fmt"
	"time"
)

type Task struct {
	Id          int
	Status      string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func main() {
	tasks := make(map[int]Task)

	newTask := Task{
		Id:          1,
		Status:      "My Project",
		Description: "Completing my first Project",
		CreatedAt:   time.Now().AddDate(2025, 03, 18),
		UpdatedAt:   time.Now(),
	}

	tasks[newTask.Id] = newTask

	fmt.Println(newTask)

	newTask, exists := tasks[1]
	if exists {
		fmt.Println("Task Found:")
	}
}
