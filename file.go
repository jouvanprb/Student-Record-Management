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