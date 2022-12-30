package main

import (
	"fmt"
)

func main() {
	var nameFirst string
	var nameLast string
	var quit bool = false

	for quit != true {
		if nameFirst == "" {
			fmt.Printf("Please enter your First Name: ")
			fmt.Scan(&nameFirst)
		}
		if nameLast == "" {
			fmt.Printf("Please enter your Last Name: ")
			fmt.Scan(&nameLast)
		}

		if (nameFirst != "") || (nameLast != "") {
			quit = true
		}
	}
	fmt.Printf("Welcome %s %s", nameFirst, nameLast)
	fmt.Println()

}
