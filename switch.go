package main

import "fmt"

func main() {
	var age int
	fmt.Printf("Enter your age:")
	fmt.Scan(&age)

	switch {
	case age < 10:
		fmt.Println("You are a child.")
	case age < 19:
		fmt.Println("You are a teen.")
	case age > 19:
		fmt.Println("You are an adult.")
	}

	var username string
	fmt.Printf("Please enter your username:")
	fmt.Scan(&username)

	switch username {
	case "john.doe@gmail.com":
		fmt.Println("Welcome!")
	default:
		fmt.Println("Username does not exist")

	}
}
