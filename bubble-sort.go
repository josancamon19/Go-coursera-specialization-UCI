package main

import (
	"fmt"
	"strconv"
)

func main() {
	var integers []int
	fmt.Println("Keep adding numbers to sort.\nPress q when you want to stop adding and sort..")
	for {
		var strInteger string
		fmt.Print("Add a number:\t")
		_, _ = fmt.Scan(&strInteger)
		if strInteger == "q" {
			break
		}
		integer, _ := strconv.Atoi(strInteger)
		integers = append(integers, integer)
	}
	BubbleSort(integers)
	fmt.Println(integers)
}

func BubbleSort(integers []int) {
	loop := true
	for loop {
		swaps := 0
		for i, _ := range integers {
			if len(integers)-1 == i {
				break
			}
			if integers[i] > integers[i+1] {
				Swap(integers, i)
				swaps += 1
			}
		}
		if swaps == 0 {
			loop = false
		}
	}
}

func Swap(integers []int, position int) {
	integers[position], integers[position+1] = integers[position+1], integers[position]
}
