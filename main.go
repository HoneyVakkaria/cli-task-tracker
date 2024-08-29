package main

import (
	"errors"
	"fmt"
	"github/honeyvakkaria/cli-task-tracker/app_actions"
	"github/honeyvakkaria/cli-task-tracker/errorhandle"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("not enough arguments")
		return
	}

	switch os.Args[1] {
	case "add":
		errorhandle.HandleLenArgsError(3)
		app_actions.AddTask(os.Args[2])
	case "update":
		errorhandle.HandleLenArgsError(4)
		app_actions.UpdateTask(convertId(os.Args[2]), os.Args[3])
	case "delete":
		errorhandle.HandleLenArgsError(3)
		app_actions.DeleteTask(convertId(os.Args[2]))
	case "list":
		status := calculateStatus()
		app_actions.List(status)
	case "mark-in-progress":
		errorhandle.HandleLenArgsError(3)
		app_actions.Mark("in-progress", convertId(os.Args[2]))
	case "mark-done":
		errorhandle.HandleLenArgsError(3)
		app_actions.Mark("done", convertId(os.Args[2]))
	default:
		fmt.Println("not supported action")
	}
}

// Function is helper to convert from argument to app_actions.ListArguments type
func calculateStatus() app_actions.ListArguments {
	if len(os.Args) == 2 {
		return app_actions.All
	}

	switch os.Args[2] {
	case "done":
		return app_actions.Done
	case "todo":
		return app_actions.Todo
	case "in-progress":
		return app_actions.Progress
	default:
		errorhandle.HandleError(errors.New("status is incorrect"))
	}

	return app_actions.All
}

// Function is helper to conver id from string to int
func convertId(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		errorhandle.HandleError(err)
	}
	return n
}
