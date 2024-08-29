package task_actions

import (
	"fmt"
	"time"
)

type Status int

const (
	Todo Status = iota
	Progress
	Done
)

type Task struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	Status      Status `json:"status"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

func CreateTask(id int, description string, status Status) Task {
	newTask := Task{
		Id:          id,
		Description: description,
		Status:      status,
		CreatedAt:   createTime(),
		UpdatedAt:   createTime(),
	}

	return newTask
}

func (t *Task) UpdateTask(description string) {
	t.Description = description
	t.UpdatedAt = createTime()
}

func (t *Task) MarkProgress() {
	t.Status = Progress
	t.UpdatedAt = createTime()
}

func (t *Task) MarkDone() {
	t.Status = Done
	t.UpdatedAt = createTime()
}

func createTime() string {
	time := time.Now()
	return fmt.Sprintf("the %dth of %s %d at %d:%d",
		time.Day(), time.Month().String(), time.Year(), time.Hour(), time.Minute())
}
