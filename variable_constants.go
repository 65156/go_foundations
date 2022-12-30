package main

import "fmt"

func main() {
	var username string
	username = "Octallium"
	fmt.Println("Username:", username)

	firstName := "Mr"
	lastName := "Giggity"
	fmt.Println("Full Name:", firstName, lastName)

	var (
		a string = "tic"
		b int    = 4
		c bool   = true
	)
	fmt.Println(a, b, c)

	var d, e, f = "i am a string", 9.0, true
	fmt.Printf("d: %v, e: %.2f, f: %v\n", d, e, f)

	const API_KEY = "dhd1345nj432ji23n5"
}
