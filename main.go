package main

import (
	"fmt"
	"github/honeyvakkaria/cli-task-tracker/json_handle"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Print("not enough arguments")
		return
	}

	name := os.Args[1]
	age, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatalf("second argument must be an integer\n%v", err)
	}

	json_handle.JsonHandle(name, age)
}
