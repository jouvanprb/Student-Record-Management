package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ShowMenu() {
	fmt.Println()
	fmt.Println(strings.Repeat("─", 35))
	fmt.Println("   🏫 STUDENT RECORD MANAGER")
	fmt.Println(strings.Repeat("─", 35))
	fmt.Println("   1.  Add Student")
	fmt.Println("   2.  View Student")
	fmt.Println("   3.  Search Student")
	fmt.Println("   4.  Update Student")
	fmt.Println("   5.  Delete Student")
	fmt.Println("   6.  Exit")
	fmt.Println(strings.Repeat("─", 35))
}

func AddStudent() {
	fmt.Println("\n👨‍🎓 Add Student")
	fmt.Println("================")
	// save func value into variable
	data := LoadStudent()

	// Bufio for read input string
	reader := bufio.NewReader(os.Stdin)

	var id int
	fmt.Print("  ID      : ")
	fmt.Scan(&id)

	fmt.Print("  Name    : ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	var age int
	fmt.Print("  Age     : ")
	fmt.Scan(&age)

	fmt.Print("  Address : ")
	address, _ := reader.ReadString('\n')
	address = strings.TrimSpace(address)

	fmt.Println()

	// make the new instance of struct
	newData := student{
		ID:      id,
		Name:    name,
		Age:     age,
		Address: address,
	}
	// add data to slice
	data = append(data, newData)
	// Send data to Func
	SaveStudent(data)
	
	fmt.Println(" ==> Data Created Successfully ")
}

func ViewStudent() {
	fmt.Println("\n📋 STUDENT LIST")
	fmt.Println("================")
	data := LoadStudent()
	for i, s := range data {
		fmt.Printf("%d. ID       : %d\n", i+1, s.ID)
		fmt.Printf("   Name     : %s\n", s.Name)
		fmt.Printf("   Age      : %d\n", s.Age)
		fmt.Printf("   Address  : %s\n", s.Address)
		fmt.Println()
	}
}

func SearchStudent() {
	fmt.Println("\n🔍 Search Student")
	fmt.Println("================")
	fmt.Print("  Enter ID: ")
	var idStudent int
	fmt.Scan(&idStudent)
	fmt.Println()

	data := LoadStudent()
	found := false
	//
	for _, s := range data {
		if idStudent == s.ID {
			fmt.Println(" ==> ID Found!")
			fmt.Println()
			fmt.Printf("  ID       : %d\n", s.ID)
			fmt.Printf("  Name     : %s\n", s.Name)
			fmt.Printf("  Age      : %d\n", s.Age)
			fmt.Printf("  Address  : %s\n", s.Address)
			found = true
			break
		}
	}
	// condition when searched data not found
	if !found {
		fmt.Printf(" ==> ID %d Not Found!", idStudent)
	}
}

func UpdateStudent() {
	fmt.Println("\n📝 Update Student")
	fmt.Println("================")
	fmt.Print("  Enter ID: ")
	var idStudent int
	fmt.Scan(&idStudent)
	fmt.Println()

	data := LoadStudent()
	reader := bufio.NewReader(os.Stdin)
	found := false

	for i, s := range data {
		if idStudent == s.ID {
			fmt.Println(" ==> ID Found!")
			fmt.Println(" ==> Update Student Data:")
			fmt.Println()

			found = true
			fmt.Printf("  ID      : %d\n", s.ID)

			fmt.Print("  Name    : ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			fmt.Print("  Age     : ")
			ageStr, _ := reader.ReadString('\n')
			ageStr = strings.TrimSpace(ageStr)

			age, err := strconv.Atoi(ageStr)
			if err != nil {
				fmt.Printf("Return Error: %v", err)
				return
			}

			fmt.Print("  Address : ")
			address, _ := reader.ReadString('\n')
			address = strings.TrimSpace(address)

			data[i].Name = name
			data[i].Age = age
			data[i].Address = address

			SaveStudent(data)
			fmt.Println()
			fmt.Println(" ==> Data Updated Successfully ")
			break
		}
	}
	if !found {
		fmt.Printf(" ==> ID %d Not Found!", idStudent)
	}
}

func DeleteStudent() {
	fmt.Println("\n📝 Delete Student")
	fmt.Println("================")
	fmt.Print("  Enter ID: ")

	var idStudent int
	fmt.Scan(&idStudent)
	fmt.Println()

	data := LoadStudent()
	found := false

	for i, s := range data {
		if idStudent == s.ID {
			data = append(data[:i], data[i+1:]...)
			found = true
			SaveStudent(data)
			fmt.Println(" ==> Data Deleted Successfully ")
			break
		}
	}

	if !found {
		fmt.Printf(" ==> ID %d Not Found!", idStudent)
	}
}
