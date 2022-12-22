package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var userInput int
	var secretNum int

	rand.Seed(time.Now().Unix())
	secretNum = rand.Intn(1000)
	fmt.Println("Secret Number:", secretNum)

	fmt.Printf("Please enter a number: ")
	fmt.Scan(&userInput)
	fmt.Println("You entered:", userInput)

	if userInput == secretNum {
		fmt.Println("You Win!!")
	} else if userInput < secretNum {
		fmt.Println("Too Low!!")
	} else if userInput > secretNum {
		fmt.Println("Too High!!")
	}
}
