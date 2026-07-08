package main

import (
	"fmt"
	"os"
	"strings"
)

func CreateFileIfNotExist(){
	var fileName string 
	fmt.Print("Enter File Name :")	
	fmt.Scan(&fileName)

	if !strings.HasSuffix(fileName, ".json") {
		fileName += ".json"
	}

	file, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("File Error %v\n", err)
		return
	} 
	
	defer file.Close()

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
	fmt.Println(students)
	return students
}
