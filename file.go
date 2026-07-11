package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func CreateFileIfNotExist() {
	// Enter file name
	const fileName = "data.json"

	// Check the file if exist
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		fmt.Println("File is not exist")
	} else {
		return
	}

	// will create file if doesn't exist
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("File Error %v\n", err)
		return
	}

	// close file
	defer file.Close()

	// automatically write file contain json format []
	_, err = file.WriteString("[]")
	if err != nil {
		fmt.Printf("Return Error %v\n", err)
		return
	}

	// Success create file message
	fmt.Printf("File %v has been created successfully \n", fileName)
}

func LoadStudent() []student {
	// Open file
	file, err := os.OpenFile("data.json", os.O_RDWR, 0644)
	if err != nil {
		fmt.Printf("Return Error: %v\n", err)
		return nil
	}
	defer file.Close()

	// Read file
	byte, err := io.ReadAll(file)
	if err != nil {
		fmt.Printf("Return Error %v\n", err)
		return nil
	}

	// Unmarshall JSON
	var students []student
	err = json.Unmarshal(byte, &students)
	if err != nil {
		fmt.Printf("Error parsing JSON %v\n", err)
		return nil
	}

	// print
	// fmt.Println(students)
	return students
}

func SaveStudent(students []student) {
	file, err := os.OpenFile("data.json", os.O_TRUNC|os.O_RDWR, 0644)
	if err != nil {
		fmt.Printf("Return Error : %v\n", err)
		return
	}
	defer file.Close()
	// convert data to json
	dataJson, err := json.MarshalIndent(students, "", "  ")
	if err != nil {
		panic(err)
	}
	// overwrite json file
	err = os.WriteFile("data.json", dataJson, 0644)
}
