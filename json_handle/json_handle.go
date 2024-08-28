package json_handle

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func CreatePerson(name string, age int) Person {
	myPerson := Person{
		Name: name,
		Age:  age,
	}

	return myPerson
}

func JsonHandle(name string, age int) {
	jsonData, err := json.MarshalIndent(CreatePerson(name, age), "", " ")
	if err != nil {
		log.Fatalf("error in creating jsonData: %v", err)
	}

	file, err := os.OpenFile("test.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("error in opening file: %v", err)
	}
	defer file.Close()

	if _, err := file.Write(jsonData); err != nil {
		log.Fatalf("error in writing data in file: %v", err)
	}
	if _, err := file.WriteString("\n"); err != nil {
		log.Fatalf("error in writing data in file: %v", err)
	}

	fmt.Println("Data were success writing!!!")
}
