package main

import "fmt"

func main() {
	name := "octy"
	age := 8
	weight := 70.45
	adult := false

	fmt.Printf("name: %T\t", name)
	fmt.Printf("age: %T\n", age)
	fmt.Printf("weight: %T\t", weight)
	fmt.Printf("adult: %T\n", adult)
}
