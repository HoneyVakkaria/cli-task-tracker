package app_actions

import (
	"fmt"
	"github/honeyvakkaria/cli-task-tracker/jsonhandle"
	"github/honeyvakkaria/cli-task-tracker/task_actions"
)

type ListArguments int

const (
	All ListArguments = iota
	Todo
	Progress
	Done
)

// Function to convert type from app_actions.ListArguments to task_actions.Status
func converStatus(arg ListArguments) task_actions.Status {
	if arg == Progress {
		return task_actions.Progress
	}
	if arg == Todo {
		return task_actions.Todo
	}

	return task_actions.Done
}

// Function to find possible max id
func findMaxId() int {
	return len(jsonhandle.ReadAllTasksFromJson())
}

// Function to create a task
func AddTask(description string) {
	newTask := task_actions.CreateTask(findMaxId(), description, task_actions.Todo)
	jsonhandle.WriteTaskToJson(append(jsonhandle.ReadAllTasksFromJson(), newTask))
	fmt.Printf("Task \"%s\" was successfully add. (ID: %d)\n", newTask.Description, newTask.Id)
}

// Function to update a task
func UpdateTask(id int, newDescription string) {
	data := jsonhandle.ReadAllTasksFromJson()
	for i := range data {
		if data[i].Id == id {
			data[i].UpdateTask(newDescription)
			jsonhandle.WriteTaskToJson(data)
			fmt.Printf("Task (ID: %d) was successfully updated\n", data[i].Id)
			return
		}
	}
	fmt.Println("Task with the given id is not exist")
}

// Function to delete a task
func DeleteTask(id int) {
	data := jsonhandle.ReadAllTasksFromJson()
	for i := range data {
		if data[i].Id == id {
			data = append(data[:i], data[i+1:]...)
			jsonhandle.WriteTaskToJson(data)
			fmt.Println("The task was successfully deleted")
			return
		}
	}
	fmt.Println("Task with the given id is not exist")
}

// Function to listing all tasks
func List(arg ListArguments) {
	status := converStatus(arg)
	var data []task_actions.Task
	for _, v := range jsonhandle.ReadAllTasksFromJson() {
		if arg == All {
			data = append(data, v)
			continue
		}
		if status == v.Status {
			data = append(data, v)
		}
	}

	for _, v := range data {
		var status string
		switch v.Status {
		case task_actions.Done:
			status = "Done"
		case task_actions.Todo:
			status = "To-Do"
		case task_actions.Progress:
			status = "In Progress"
		}
		fmt.Printf("Task \"%s\" (ID: %d) Status: %s\n Created at: %s\n Last update was on: %s\n\n",
			v.Description, v.Id, status, v.CreatedAt, v.UpdatedAt)
	}
}

// Function to mark a task as "done" or as "in-progress"
func Mark(arg string, id int) {
	data := jsonhandle.ReadAllTasksFromJson()
	for i := range data {
		if data[i].Id == id {
			if arg == "done" {
				data[i].MarkDone()
			} else {
				data[i].MarkProgress()
			}
			jsonhandle.WriteTaskToJson(data)
			fmt.Printf("Task \"%s\" successfully marked as %s\n", data[i].Description, arg)
			return
		}
		fmt.Println("Task with the given id is not exist")
	}
}
