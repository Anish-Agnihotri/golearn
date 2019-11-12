package main

import (
	"fmt"
	"math"
)

const myConstantString string = "Some constant string"
const myConstantSin float64 = 35

func main() {
	fmt.Println(myConstantString)
	fmt.Println(math.Sin(myConstantSin))
}
