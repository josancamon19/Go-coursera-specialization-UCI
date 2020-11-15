package main

import "fmt"

type Animal struct {
	Name             string
	FoodEaten        string
	LocomotionMethod string
	SpokenSound      string
}

func (animal *Animal) eat() {
	fmt.Printf("Eating %s", animal.FoodEaten)
}

func (animal *Animal) move() {
	fmt.Printf("Moving: %s", animal.LocomotionMethod)
}

func (animal *Animal) speak() {
	fmt.Printf("%s", animal.SpokenSound)
}

func main() {
	cow := Animal{"cow", "grass", "walk", "moo"}
	bird := Animal{"bird", "worms", "fly", "peep"}
	snake := Animal{"snake", "mice", "slither", "hsss"}

	for {
		var animalName string
		fmt.Print("Type one animal name\nAvailable animals: cow, bird, snake: \t")
		_, _ = fmt.Scan(&animalName)

		var animal Animal
		if animalName == "cow" {
			animal = cow
		} else if animalName == "bird" {
			animal = bird
		} else if animalName == "snake" {
			animal = snake
		} else {
			fmt.Println("Please type a valid animal")
			continue
		}

		var action string
		fmt.Print("Type one action\nAvailable actions: eat, move, speak: \t")
		_, _ = fmt.Scan(&action)

		if action == "eat" {
			animal.eat()
		} else if action == "move" {
			animal.move()
		} else if action == "speak" {
			animal.speak()
		} else {
			fmt.Println("Try again and type a valid action")
			continue
		}
		fmt.Println("...")

	}
}
