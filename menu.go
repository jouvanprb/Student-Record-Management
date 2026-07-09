package main

import (
	"fmt"
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
