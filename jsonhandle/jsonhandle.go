package jsonhandle

import (
	"encoding/json"
	"github/honeyvakkaria/cli-task-tracker/errorhandle"
	"github/honeyvakkaria/cli-task-tracker/task_actions"
	"os"
)

// Function to add new task (or to create the first one)
// This function not append but rewriting a data
func WriteTaskToJson(data []task_actions.Task) {
	jsonData, err := json.MarshalIndent(data, "", " ")
	errorhandle.HandleError(err)

	file, err := os.OpenFile("data.json", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	errorhandle.HandleError(err)
	defer file.Close()

	_, err = file.Write(jsonData)
	errorhandle.HandleError(err)
}

// Function to read all data from "data.json" file
func ReadAllTasksFromJson() []task_actions.Task {
	var tasks []task_actions.Task
	file, err := os.Open("data.json")
	if err != nil {
		if os.IsNotExist(err) {
			return tasks
		}
		errorhandle.HandleError(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tasks)
	errorhandle.HandleError(err)

	return tasks
}
