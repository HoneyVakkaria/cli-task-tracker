package jsonhandle

import (
	"encoding/json"
	"github/honeyvakkaria/cli-task-tracker/task_actions"
	"log"
	"os"
)

// Function to add new task (or to create the first one)
// This function not append but rewriting a data
func WriteTaskToJson(data []task_actions.Task) {
	jsonData, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Fatalf("error in creating jsonData in WriteTaskToJson func:\n%v", err)
	}

	file, err := os.OpenFile("data.json", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatalf("error in opening (or creating) file in WriteTaskToJson func:\n%v", err)
	}
	defer file.Close()

	if _, err := file.Write(jsonData); err != nil {
		log.Fatalf("error in writing jsonData to file in WriteTaskToJson func:\n%v", err)
	}
}

// Function to read all data from "data.json" file
func ReadAllTasksFromJson() []task_actions.Task {
	var tasks []task_actions.Task
	file, err := os.Open("data.json")
	if err != nil {
		if os.IsNotExist(err) {
			return tasks
		}
		log.Fatalf("error in opening file in readAllTasksFromJson func:\n%v", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&tasks); err != nil {
		log.Fatalf("error in reading file in readAllTasksFromJson func: %v", err)
	}

	return tasks
}
