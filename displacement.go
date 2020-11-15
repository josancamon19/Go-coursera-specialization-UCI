package main

import "fmt"

func main() {
	var acceleration, velocity0, displacement0 float64

	fmt.Print("Enter a value for acceleration:\t")
	_, _ = fmt.Scan(&acceleration)
	fmt.Print("Enter a value for velocity:\t")
	_, _ = fmt.Scan(&velocity0)
	fmt.Print("Enter a value for displacement:\t")
	_, _ = fmt.Scan(&displacement0)

	displacementFunction := GenDisplaceFunc(acceleration, velocity0, displacement0)

	var time float64
	fmt.Print("Enter a value for time (seconds):\t")
	_, _ = fmt.Scan(&time)

	fmt.Printf("Displacement after %v seconds is %v", time, displacementFunction(time))
}

func GenDisplaceFunc(acceleration, velocity0, displacement0 float64) func(time float64) float64 {
	return func(time float64) float64 {
		return (0.5 * acceleration * time * time) + (velocity0 * time) + displacement0
	}
}
