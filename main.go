package main

import (
	"fmt"
	"strings"
)

func main() {
	CreateFileIfNotExist()

	for {
		ShowMenu()

		var choice int
		fmt.Print("Choose menu (1-6): ")
		fmt.Scan(&choice)
		fmt.Println(strings.Repeat("─", 35))

		switch choice {
		case 1:
			AddStudent()
		case 2:
			ViewStudent()
		case 6:
			fmt.Println("-> PROGRAM STOPPED!")
			fmt.Println()
			return
		}
	}

}
