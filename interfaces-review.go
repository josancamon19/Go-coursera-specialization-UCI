package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type AnimalReview2 interface {
	Eat()
	Speak()
	Move()
}

type CowReview struct {
	food       string
	locomotion string
	noise      string
}

func (a CowReview) Eat() {
	fmt.Println(a.food)
}

func (a CowReview) Speak() {
	fmt.Println(a.noise)
}

func (a CowReview) Move() {
	fmt.Println(a.locomotion)
}

type BirdReview struct {
	food       string
	locomotion string
	noise      string
}

func (a BirdReview) Eat() {
	fmt.Println(a.food)
}

func (a BirdReview) Speak() {
	fmt.Println(a.noise)
}

func (a BirdReview) Move() {
	fmt.Println(a.locomotion)
}

type SnakeReview struct {
	food       string
	locomotion string
	noise      string
}

func (a SnakeReview) Eat() {
	fmt.Println(a.food)
}

func (a SnakeReview) Speak() {
	fmt.Println(a.noise)
}

func (a SnakeReview) Move() {
	fmt.Println(a.locomotion)
}

func main() {
	dict := make(map[string]AnimalReview2)
	animals := map[string]AnimalReview2{
		"cow": CowReview{"grass", "walk", "moo"},
		"bird": BirdReview{
			"worms", "fly", "peep",
		},
		"snake": SnakeReview{
			"mice", "slither", "hsss",
		},
	}
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		cmdString, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		strings.TrimSuffix(cmdString, "\n")
		splitted := strings.Fields(cmdString)
		if splitted[0] == "exit" {
			fmt.Println("quiting...bye...bye")
			return
		}
		err = validateOperations(splitted)
		if err != nil {
			fmt.Println(err)
			continue
		}
		name := splitted[1]

		switch splitted[0] {
		case "newanimal":
			newAnimal := splitted[2]
			if err := validateAnimals(newAnimal); err != nil {
				fmt.Println(err)
				continue
			}
			dict[name] = animals[splitted[2]]
			fmt.Println("Created it!")
		case "query":
			err := validateActions(splitted)

			if err != nil {
				fmt.Println(err)
				continue
			}

			operation, ok := dict[name]
			if !ok {
				fmt.Printf("name %s not found, please run newanimal first\n", name)
				continue
			}

			switch splitted[2] {
			case "eat":
				operation.Eat()
			case "move":
				operation.Move()
			case "speak":
				operation.Speak()
			}
		}
	}
}

func validateOperations(userInput []string) error {
	validOperations := map[string]bool{
		"query":     true,
		"newanimal": true,
	}
	if len(userInput) < 3 {
		return errors.New("wrong number of arguments. usage: newanimal name cow")
	}
	if !validOperations[userInput[0]] {
		return errors.New("wrong first argument, valid: 'query' or 'newanimal' without quotes")

	}
	return nil
}

func validateActions(userInput []string) error {
	validActions := map[string]bool{
		"eat":   true,
		"speak": true,
		"move":  true,
	}
	if !validActions[userInput[2]] {
		return errors.New("wrong operation, valid: 'eat' or 'speak' or 'move' without quotes")

	}
	return nil
}

func validateAnimals(animal string) error {
	validAnimals := map[string]bool{
		"cow":   true,
		"bird":  true,
		"snake": true,
	}
	if !validAnimals[animal] {
		return errors.New("wrong animal, please use one of: cow, bird or snake")
	}
	return nil
}
