package errorhandle

import (
	"errors"
	"fmt"
	"log"
	"os"
	"runtime"
)

// TODO! Convert any error to custom error
// Function to handle any error
func HandleError(err error) {
	if err == nil {
		return
	}

	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		fmt.Println("Can't read an info of error")
		return
	}

	fn := runtime.FuncForPC(pc).Name()
	log.Fatalf("error: %v\nfrom file: %s\nfrom function: %s\nfrom line: %d", err, file, fn, line)
}

// Function to handle only errors from main function
func HandleLenArgsError(n int) {
	if len(os.Args) < n {
		err := errors.New("not enough arguments")
		log.Fatal(err)
	}
}
