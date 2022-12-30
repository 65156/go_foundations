package main

import "fmt"

func main() {
	// if - else
	num := 10
	if num < 10 {
		fmt.Println("Less than 10.")
	} else if num > 10 {
		fmt.Println("More than 10.")
	} else if num == 10 {
		fmt.Println("Equal to 10.")
	}

	username := "octallium"
	age := 14
	usernamecheck := false
	username = "Octallium"
	if username == "Octallium" || username == "octallium" {
		usernamecheck = true
	} else {
		fmt.Println("Wrong Username.")
	}

	if usernamecheck == true && age == 14 {
		fmt.Println("Passed Verification")
	}
}
