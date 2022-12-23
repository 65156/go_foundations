package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var userInput int
	var secretNum int
	var quit bool = false

	rand.Seed(time.Now().Unix())
	secretNum = rand.Intn(100)
	fmt.Println("Secret Number:", secretNum)

	for quit != true {
		fmt.Printf("Please enter a number: ")
		fmt.Scan(&userInput)

		if userInput == secretNum {
			fmt.Println("You Win!!")
			quit = true
		} else if userInput < secretNum {
			fmt.Println("Too Low...")
		} else if userInput > secretNum {
			fmt.Println("Too High...")
		}
	}

}
