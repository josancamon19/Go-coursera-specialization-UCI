package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Type a float\n")
	var input float64
	fmt.Scan(&input)
	fmt.Printf("%.0f", input)
}