package main

import (
	"fmt"
	"math"
)

func main() {
	var1 := math.Pi
	var2 := "Hey Gropher!"
	var3 := true
	var4 := 49

	fmt.Printf("Value: %v\tData Type: %T\n", var1, var1)
	fmt.Printf("Value: %v\tData Type: %T\n", var2, var2)
	fmt.Printf("Value: %v\tData Type: %T\n", var3, var3)
	fmt.Printf("Value: %v\tData Type: %T\n", var4, var4)

	var5 := false

	if var5 == false {
		var1 = 123.5324
		var2 = "James Bond"
		var3 = true
		var4 = 3672 - 982 + 1293
		var5 = true

		fmt.Printf("Value: %v\tData Type: %T\n", var1, var1)
		fmt.Printf("Value: %v\tData Type: %T\n", var2, var2)
		fmt.Printf("Value: %v\tData Type: %T\n", var3, var3)
		fmt.Printf("Value: %v\tData Type: %T\n", var4, var4)
		fmt.Printf("Value: %v\tData Type: %T\n", var5, var5)

	}
}
