package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// Creating a struct for my Task

type Task struct {
	Id          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

var tasks []Task

func main() {
	const fileName = "Task.json"
	if err := loadTasks(fileName); err != nil {
		fmt.Println("Error Loading Task: ", err)
		return
	}
	fmt.Println("Welcome to Task Tracker")

	if len(os.Args) < 3 {
		fmt.Println("Usage: task-tracker-cli [command] [args]")
		return
	}
	description := os.Args[2]

	newTask := Task{
		len(tasks) + 1,
		description,
		"not_done",
		time.Now(),
		time.Now(),
	}
	tasks = append(tasks, newTask)
}

func loadTasks(fileName string) error {
	file, err := os.ReadFile(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	return json.Unmarshal(file, &tasks)
}

func saveToTask(fileName string) error {
	data, err := json.MarshalIndent(tasks, "", "")
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, data, 0644)
}

func updateTask() {

}
