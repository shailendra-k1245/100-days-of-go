package main

import (
	"fmt"
)

func main() {
	// example of signed integer
	var x int = 500
	var y int = -4500

	// unsigned integer
	var i unit = 600
	var j unit = 4600

	fmt.Printf("Type : %T, value: %v \n", x, x)
	fmt.Printf("Type: %T, value: %v \n", y, y)

	fmt.Printf("Type: %T, value: %v", x, x)
	fmt.Printf("Type: %T, value: %v", y, y)
}
