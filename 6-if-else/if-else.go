package main

import "fmt"

func main() {
	// FizzBuzz
	for i := 1; i <= 100; i++ {
		if i%15 == 0 {
			// If multiple of 15, print FizzBuzz (could also be written i%3 == 0 && i%5 == 0)
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			// If multiple of 3, print Fizz
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			// If multiple of 5, print Buzz
			fmt.Println("Buzz")
		} else {
			// If anything else, print the number
			fmt.Println(i)
		}
	}
}
