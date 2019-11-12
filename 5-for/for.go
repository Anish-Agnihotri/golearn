package main

import "fmt"

func main() {
	// Value print loop
	for i := 0; i <= 5; i++ {
		fmt.Println("Value of i is:", i)
	}
	// Infinite loop
	for {
		fmt.Println("infinite")
		break
	}
	// Conditional
	for i := 0; i <= 5; i++ {
		if i%2 == 0 {
			fmt.Println("2, yes", i)
			continue
		}
		fmt.Println(i)
	}
}
