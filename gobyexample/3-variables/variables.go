package main

import "fmt"

func main() {
	fmt.Println("Testing declaring strings, integers, and booleans:")

	var myName string = "Anish"
	var myAge int = 17
	var amISick bool = true

	fmt.Println("My name is ", myName)
	fmt.Println("My age is ", myAge)
	fmt.Println("Am I sick? ", amISick)

	shorthand := "This is shorthand decleration"
	fmt.Println(shorthand)
}
