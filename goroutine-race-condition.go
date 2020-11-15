package main

import (
	"fmt"
	"gorm.io/gorm/utils"
	"time"
)

func main() {
	// Assertion will be faster than the increase
	// I'd say as Increase has to write to a pointer, and assert only reads

	value := 0
	go Increase(&value)
	go AssertValue(&value, 1)
	time.Sleep(time.Second) // Otherwise the go routines would not have time to initiate

	fmt.Printf("Final integer result:\t %d", value)
}

func Increase(integer *int) {
	*integer += 1
}

func AssertValue(integer *int, expectedValue int) {
	fmt.Printf("Integer was increased:\t %t \n", utils.AssertEqual(integer, expectedValue))
}
