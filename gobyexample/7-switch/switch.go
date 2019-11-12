package main

import (
	"fmt"
	"time"
)

func main() {
	// Testing basic switch
	for i := 1; i <= 3; i++ {
		switch i {
		case 1:
			fmt.Println("3")
		case 2:
			fmt.Println("2")
		case 3:
			fmt.Println("1")
		}
	}

	// Testing switch usage with time
	date := time.Now().Weekday()
	switch date {
	case time.Saturday:
		fmt.Println("It's saturday, woo!")
	case time.Tuesday:
		fmt.Println("It's tuesday :(")
	}

	// Testing function with switch usage
	whatAmI := func(i interface{}) {
		switch i.(type) {
		case string:
			fmt.Println("I'm a string")
		case int:
			fmt.Println("I'm an integer")
		default:
			fmt.Println("Who knows what this is! Float them!")
		}
	}

	whatAmI("Anish is cool") // string
	whatAmI(11)              // integer
	whatAmI(11.0)            // float
}
