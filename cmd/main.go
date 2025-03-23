package main

import (
	"fmt"
	"os"
	"strconv"
	"task-cli/internal/storage"
	"task-cli/internal/task"
	"time"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("No input")
		return
	}

	storageFileName := "tasks.json"
	storage := storage.JsonStorage{
		JsonFilePath: storageFileName,
	}
	storage.LoadStorageFromJsonFile()

	command := args[1]
	switch command {
	case "add":
		taskDescription := args[2]
		newTask := task.Task{
			Description: taskDescription,
			Status:      task.TODO,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}
		taskId, _ := storage.AddTask(newTask)
		fmt.Printf("Task added successfully (ID: %d)", taskId)
	case "update":
		taskId, _ := strconv.Atoi(args[2])
		taskDescription := args[3]
		thisTask := storage.GetTaskById(taskId)
		thisTask.Description = taskDescription
		thisTask.UpdatedAt = time.Now()
		storage.UpdateTask(thisTask)
	case "delete":
		taskId, _ := strconv.Atoi(args[2])
		storage.DeleteTask(taskId)
	case "mark-in-progress", "mark-done":
		taskId, _ := strconv.Atoi(args[2])
		thisTask := storage.GetTaskById(taskId)
		thisTask.UpdatedAt = time.Now()
		if command == "mark-in-progress" {
			thisTask.Status = task.IN_PROGRESS
		} else {
			thisTask.Status = task.DONE
		}
		storage.UpdateTask(thisTask)
	case "list":
		statusFilter := ""
		if len(args) == 3 {
			statusFilter = args[2]
		}
		storage.ListTasksByStatus(statusFilter)
	default:
		fmt.Printf("command: `%s` not supported\n", command)
		return
	}
}
