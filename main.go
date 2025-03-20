package main

import (
	"encoding/json"
	"fmt"
	"os"
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

func main() {

	const fileName = "tasks.json"

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
			Id:          len(tasks) + 1,
			Description: description,
			Status:      "not_done",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
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
		if err != nil {
			fmt.Println("Invalid ID:", err)
			return
		}
		description := os.Args[3]
		if err := updateTask(id, description); err != nil {
			fmt.Println("Error Updating Task", err)
			return
		}
		fmt.Println("Task updated successfully.")

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: task-tracker-cli delete [id]")
			return
		}
		idStr := os.Args[2]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Println("Invalid ID:", err)
			return
		}
		if err := deleteTask(id); err != nil {
			fmt.Println("Error deleting task:", err)
			return
		}
		fmt.Println("Task deleted successfully.")

	case "Mark-In-Progress":
		if len(os.Args) < 3 {
			fmt.Println("Usage: task-tracker-cli mark-done [id]")
			return
		}
		idStr := os.Args[2]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Println("Invalid ID:", err)
			return
		}
		if err := markInProgress(id); err != nil {
			fmt.Println("Error marking task Progress:", err)
			return
		}
		fmt.Println("Marked Progress Task")

	case "MarkDone":
		if len(os.Args) < 3 {
			fmt.Println("Usage: task-tracker-cli mark-done [id]")
			return
		}
		idStr := os.Args[2]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Println("Invalid ID:", err)
			return
		}
		if err := markDone(id); err != nil {
			fmt.Println("Task Marked Done")
		}

	case "list":
		var status string
		if len(os.Args) < 2 {
			status = os.Args[2]
		}
		taskList(status)

	default:
		fmt.Println("Unknown command:", command)
	}

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
	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, data, 0644)
}

func updateTask(id int, description string) error {
	for i, task := range tasks {
		if task.Id == id {
			tasks[i].Description = description
			tasks[i].UpdatedAt = time.Now()
			return saveToTask("tasks.json")

		}
	}
	return fmt.Errorf("task with ID %d not found", id)
}

func deleteTask(id int) error {
	for i, task := range tasks {
		if task.Id == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return saveToTask("tasks.json")
		}
	}
	return fmt.Errorf("task with ID %d not found", id)
}

func markInProgress(id int) error {
	for i, task := range tasks {
		if task.Id == id {
			tasks[i].Status = "In-Progress"
			tasks[i].UpdatedAt = time.Now()
			return saveToTask("tasks.json")
		}
	}
	return fmt.Errorf("task with ID %d not found", id)
}

func markDone(id int) error {
	for i, task := range tasks {
		if task.Id == id {
			tasks[i].Status = "Done"
			tasks[i].UpdatedAt = time.Now()
			return saveToTask("tasks.json")
		}
	}
	return fmt.Errorf("task with ID %d not found", id)
}

func taskList(status string) {
	for _, task := range tasks {
		if status == "" || task.Status == status {
			fmt.Printf("ID: %d, Description: %s, Status: %s, CreatedAt: %s, UpdatedAt: %s\n",
				task.Id,
				task.Description,
				task.Status,
				task.CreatedAt.Format("20025-03-19 15:52:10"),
				task.UpdatedAt.Format("2025-03-19 15:52:10"))
		}
	}
}
