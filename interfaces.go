package main

import (
	"fmt"
)

type NewAnimal interface {
	Speak()
	Move()
	Eat()
}

type Snake struct {
	Name string
}

func (animal Snake) Speak() {
	fmt.Printf("SnakeReview %s says hsss\n", animal.Name)
}

func (animal Snake) Eat() {
	fmt.Printf("SnakeReview %s eats mice\n", animal.Name)
}

func (animal Snake) Move() {
	fmt.Printf("SnakeReview %s slither\n", animal.Name)
}

type Cow struct {
	Name string
}

func (animal Cow) Speak() {
	fmt.Printf("CowReview %s says moo\n", animal.Name)
}

func (animal Cow) Eat() {
	fmt.Printf("CowReview %s eats grass\n", animal.Name)
}

func (animal Cow) Move() {
	fmt.Printf("CowReview %s walks\n", animal.Name)
}

type Bird struct {
	Name string
}

func (animal Bird) Speak() {
	fmt.Printf("BirdReview %s says BirdReview\n", animal.Name)
}

func (animal Bird) Eat() {
	fmt.Printf("BirdReview %s eats worms\n", animal.Name)
}

func (animal Bird) Move() {
	fmt.Printf("BirdReview %s fly\n", animal.Name)
}

func main() {

	var animals []NewAnimal
	for {
		fmt.Print("Type a command: \t")
		var command, command1, command2 string
		_, _ = fmt.Scan(&command, &command1, &command2)
		commands := []string{command, command1, command2}
		if len(commands) != 3 {
			fmt.Println("Invalid command, try again ...")
			continue
		}

		operation := commands[0]
		if operation == "newanimal" {
			AppendAnimal(commands, &animals)
			fmt.Println(animals)
		} else if operation == "query" {
			QueryAnimalAction(commands, animals)
		}
	}
}

func AppendAnimal(commands []string, animals *[]NewAnimal) {
	animalType := commands[2]
	if animalType == "snake" {
		*animals = append(*animals, Snake{commands[1]})
	} else if animalType == "cow" {
		*animals = append(*animals, Cow{commands[1]})
	} else if animalType == "bird" {
		*animals = append(*animals, Bird{commands[1]})
	}
}

func QueryAnimalAction(commands []string, animals []NewAnimal) {
	name := commands[1]
	action := commands[2]

	for _, animalInterface := range animals {
		var animal NewAnimal

		snake, ok := animalInterface.(Snake)
		if ok && snake.Name == name {
			animal = snake
		}

		cow, ok := animalInterface.(Cow)
		if ok && cow.Name == name {
			animal = cow
		}

		bird, ok := animalInterface.(Bird)
		if ok && bird.Name == name {
			animal = bird
		}

		if animal != nil {
			if action == "eat" {
				animal.Eat()
			} else if action == "move" {
				animal.Move()
			} else if action == "speak" {
				animal.Speak()
			}
			break
		}

	}
}
