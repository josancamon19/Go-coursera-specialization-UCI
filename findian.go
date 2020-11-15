package main

import (
	"fmt"
	"strings"
)

func main() {
	var text string
	fmt.Print("Enter a string: \t")
	_, err := fmt.Scan(&text)
	if err != nil {
		panic(err)
	}
	text = strings.ToLower(text)
	if text[0] == 'i' && text[len(text)-1] == 'n' && strings.Contains(text, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not found!")
	}
}
