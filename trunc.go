package main

import "fmt"

func main() {
	var floatingPoint float64
	fmt.Print("Enter a floating number: \t")
	_, err := fmt.Scan(&floatingPoint)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Truncated version %.0f \t Floating point %v", floatingPoint, floatingPoint)
}
