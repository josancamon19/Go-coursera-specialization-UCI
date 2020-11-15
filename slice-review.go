package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// prompt to get integer
	var theIntAsStr string
	var intSlice = make([]int, 3)

	// loop until an x or X is entered
	for i := 0; strings.Compare(strings.ToLower(theIntAsStr), "x") != 0; i++ {
		fmt.Printf("Enter an integer or press 'x' to finish : ")
		fmt.Scan(&theIntAsStr)
		// check entered character is a number and add to array if it is
		if num, err := strconv.Atoi(theIntAsStr); err == nil {
			// prevent overwriting due to sorting in first three iterations
			// is there a better way?
			if i < 3 {
				ii := sort.SearchInts(intSlice[0:3], 0)
				intSlice[ii] = num
			} else {
				intSlice = append(intSlice, num)
			}
			// sort and print
			sort.Ints(intSlice)
			fmt.Println(intSlice)
		}
	}

}
