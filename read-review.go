package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"strings"
)

func main() {

	type Person struct {
		fname string
		lname string
	}

	//Retrieve file name
	fmt.Printf("Please enter the file name:")
	consoleReader := bufio.NewReader(os.Stdin)
	fileName, err := consoleReader.ReadString('\n')
	//Remove \r\n for the string enter by the user
	if runtime.GOOS == "windows" {
		fileName = strings.TrimRight(fileName, "\r\n")
	} else {
		fileName = strings.TrimRight(fileName, "\n")
	}

	// Open and Read the file
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		print("Error when opening this file... Not found\n")
	}
	splitString := strings.Fields(string(data))

	// Creating Slice
	slice := []Person{}

	// Add Each Struct in slice with only 20 characters
	for i := 0; i < len(splitString); i = i + 2 {
		fnameLine := splitString[i]
		if len(fnameLine) > 20 {
			fnameLine := fnameLine[0:19]
			_ = fnameLine
		}
		lnameLine := splitString[i+1]
		if len(lnameLine) > 20 {
			lnameLine := lnameLine[0:19]
			_ = lnameLine
		}
		person := Person{fname: fnameLine, lname: lnameLine}
		slice = append(slice, person)
	}

	for i, v := range slice {
		fmt.Printf("Struct %d - fname : %s - lname : %s\n", i, v.fname, v.lname)
	}

}
