package main

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime/metrics"
	"strconv"
	"time"
)

// Creating a struct for my Task

type Task struct {
	Id          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"` // "Not_Done", "In-progress", "Done"
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

var tasks []Task

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

func updateTask(id int, description string) {
	for i, tasks := range tasks {
		if tasks.Id == id {

		}
	}
}

func main() {
	const fileName = "Task.json"
	if err := loadTasks(fileName); err != nil {
		fmt.Println("Error Loading Task: ", err)
		return
	}
	fmt.Println("Welcome to Task Tracker")

	if len(os.Args) < 2 {
		fmt.Println("Usage: task-tracker-cli [command] [args]")
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Please pick a Task description.")
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

		if err := saveToTask(fileName); err != nil {
			fmt.Println("Error Saving Task:", err)
			return
		}
		fmt.Printf("Task Added: ID: %d, Description:%s, Status:%s, CreatedAt:%s, UpdatedAt:%s\n",
			newTask.Id,
			newTask.Description,
			newTask.Status,
			newTask.CreatedAt.Format("2025-03-19 15:52:10"),
			newTask.UpdatedAt.Format("2025-03-19 15:52:10"),
		)
	case "update":
		if len(os.Args) < 4 {
			fmt.Println("Usage: task-tracker-cli update [id] [description]")
			return
		}
		idStr := os.Args[2]
		id, err := strconv.Atoi(idStr)
	}

}
