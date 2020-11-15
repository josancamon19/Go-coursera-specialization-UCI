package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const prefix = "i"
const middle = "a"
const suffix = "n"

func main() {

	fmt.Printf("go ahead and enter your string\n")

	var scanner = bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var input = scanner.Text()

	input = strings.ToLower(input)

	if strings.HasPrefix(input, prefix) && strings.HasSuffix(input, suffix) && strings.Contains(input, middle) {
		fmt.Printf("Found!")
	} else {
		fmt.Printf("Not found!")
	}
}
