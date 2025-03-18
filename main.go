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

	tasks := []Task{
		{1,
			"My first Project",
			"Completing my first Project",
			time.Now().AddDate(2025, 03, 18),
			time.Now(),
		},
	}
	fmt.Println(tasks)

	newTask := Task{
		Id:          2,
		Status:      "My second Project",
		Description: "Completing my Second Project",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	tasks = append(tasks, newTask)
	fmt.Println(tasks)

}
