package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func main() {
	fmt.Print("Enter the name of your txt file:\t")
	var fileName string
	_, _ = fmt.Scan(&fileName)

	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	var persons []Person
	for _, fullName := range strings.Split(string(content), "\n") {
		nameParts := strings.Split(fullName, " ")
		if len(nameParts) != 2 {
			fmt.Println("Invalid name, E.g.: Joan Cabezas")
			continue
		}
		persons = append(persons, Person{
			FirstName: nameParts[0],
			LastName:  nameParts[1],
		})
	}
	fmt.Println(persons)

}
