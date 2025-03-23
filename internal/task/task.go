package task

import (
	"time"
)

type Task struct {
	TaskID      int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

var TODO = "todo"
var IN_PROGRESS = "in-progress"
var DONE = "done"
