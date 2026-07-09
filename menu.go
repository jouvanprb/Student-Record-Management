package main

import (
	"bufio"
	"fmt"
	"os"
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
	fmt.Print("ID     \n: ")
	fmt.Scan(&id)

	fmt.Print("Name   \n: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	var age int
	fmt.Print("Age    \n: ")
	fmt.Scan(&age)

	fmt.Print("Address\n: ")
	address, _ := reader.ReadString('\n')
	address = strings.TrimSpace(address)

	// make the new instance of struct
	newData := student{
		ID:      id,
		Name:    name,
		Age:     age,
		Address: address,
	}
	// add data to slice
	data = append(data, newData)

	SaveStudent(data)
}

func ViewStudent() {
	fmt.Println("\n📋 STUDENT LIST")
	fmt.Println("================")
	data := LoadStudent()
	for i, s := range data {
		fmt.Printf("%d ID       : %d\n", i+1, s.ID)
		fmt.Printf("  Name     : %s\n", s.Name)
		fmt.Printf("  Age      : %d\n", s.Age)
		fmt.Printf("  Address  : %s\n", s.Address)
		fmt.Println()
	}
}
