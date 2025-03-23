package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"task-cli/internal/task"
)

type JsonStorage struct {
	JsonFilePath string
	taskDict     map[int]task.Task
}

func (js *JsonStorage) saveStorageToJsonFile() error {
	taskList := js.parseJsonStorageToTaskList()
	jsonString, err := json.Marshal(taskList)
	if err != nil {
		return err
	}
	err = os.WriteFile(js.JsonFilePath, jsonString, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (js *JsonStorage) parseTaskListToJsonStorage(taskList []task.Task) error {
	js.taskDict = make(map[int]task.Task)
	for _, task := range taskList {
		js.taskDict[task.TaskID] = task
	}
	return nil
}

func (js *JsonStorage) parseJsonStorageToTaskList() []task.Task {
	var taskList []task.Task
	for _, v := range js.taskDict {
		taskList = append(taskList, v)
	}
	return taskList
}

func (js *JsonStorage) LoadStorageFromJsonFile() error {
	var taskList []task.Task
	// Read storage from json file
	f, err := os.ReadFile(js.JsonFilePath)
	if err == nil {
		if len(f) > 0 {
			err = json.Unmarshal(f, &taskList)
			if err != nil {
				log.Fatal(err)
			}
		}
	} else if errors.Is(err, os.ErrNotExist) {
		ff, err := os.OpenFile(js.JsonFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		ff.Close()
	} else {
		return err
	}
	return js.parseTaskListToJsonStorage(taskList)
}

// Task related methods
func (js *JsonStorage) GetTaskById(taskId int) task.Task {
	return js.taskDict[taskId]
}
func (js *JsonStorage) getNewId() int {
	maxId := 0
	for id, _ := range js.taskDict {
		if id > maxId {
			maxId = id
		}
	}
	return maxId + 1
}

func (js *JsonStorage) AddTask(task task.Task) (int, error) {
	task.TaskID = js.getNewId()
	fmt.Println(task.TaskID)
	js.taskDict[task.TaskID] = task

	fmt.Println("Save Tasks to Storage")
	js.saveStorageToJsonFile()
	return task.TaskID, nil
}

func (js *JsonStorage) UpdateTask(task task.Task) error {
	js.taskDict[task.TaskID] = task
	js.saveStorageToJsonFile()
	return nil
}

func (js *JsonStorage) DeleteTask(taskID int) {
	delete(js.taskDict, taskID)
	js.saveStorageToJsonFile()
}

func (js *JsonStorage) ListTasksByStatus(statusFilter string) {
	for _, v := range js.taskDict {
		if statusFilter == "" {
			fmt.Println(v)
		} else {
			if v.Status == statusFilter {
				fmt.Println(v)
			}
		}
	}
}
