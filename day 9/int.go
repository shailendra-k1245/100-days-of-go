package main

import (
	"fmt"
)

func main() {
	// example of signed integer
	var x int = 500
	var y int = -4500

	// unsigned integer
	var i uint = 600
	var j uint = 4600

	fmt.Printf("Type : %T, value: %v \n", x, x)
	fmt.Printf("Type: %T, value: %v \n", y, y)

	fmt.Printf("Type: %T, value: %v \n", i, i)
	fmt.Printf("Type: %T, value: %v", j, j)
}
