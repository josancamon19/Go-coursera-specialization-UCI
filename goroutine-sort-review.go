package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	var s string
	var lenArr int
	channel := make(chan []int)
	scanner := bufio.NewScanner(os.Stdin)
	finalArr := make([]int, 0)
	fmt.Println("Please enter as many numbers separated by the space as you want in one line:  ")
	scanner.Scan()
	s = strings.TrimSpace(scanner.Text())
	arrNum := strings.Split(s, " ")
	lenArr = len(arrNum) / 4

	go toInt(arrNum, channel, lenArr, 0)
	go toInt(arrNum, channel, lenArr, lenArr)
	go toInt(arrNum, channel, lenArr, lenArr*2)
	go toInt(arrNum, channel, 0, lenArr*3)

	finalArr = append(finalArr, <-channel...)
	finalArr = append(finalArr, <-channel...)
	finalArr = append(finalArr, <-channel...)
	finalArr = append(finalArr, <-channel...)

	sort.Ints(finalArr)
	fmt.Println(finalArr)
}

func toInt(sli []string, c chan []int, len int, startFrom int) {
	intArr := make([]int, 0)
	var l int
	if len == 0 {
		l = cap(sli)
	} else {
		l = startFrom + len
	}

	for i := startFrom; i < l; i++ {
		n, err := strconv.Atoi(sli[i])
		if err != nil {
			fmt.Printf("error with element %s, error %s: ", sli[i], err.Error())
		}
		intArr = append(intArr, n)
	}
	sortArray(intArr, c)
}

func sortArray(sli []int, c chan []int) {
	sort.Ints(sli)
	fmt.Println("Part of array sorted in Goroutine: ", sli)
	c <- sli
}
