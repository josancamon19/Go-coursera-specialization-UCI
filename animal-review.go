package main

import "fmt"

type AnimalReview struct {
	Food       string
	Locomotion string
	Sound      string
}

func (a *AnimalReview) Eat() {
	fmt.Println(a.Food)
}

func (a *AnimalReview) Move() {
	fmt.Println(a.Locomotion)
}

func (a *AnimalReview) Speak() {
	fmt.Println(a.Sound)
}

func main() {
	// Initilize animal
	var animal, info string

	validAnimal := []string{"cow", "bird", "snake"}
	validAnimalInfo := []string{"eat", "move", "speak"}

	cow := AnimalReview{Food: "grass", Locomotion: "walk", Sound: "moo"}
	bird := AnimalReview{Food: "worms", Locomotion: "fly", Sound: "peep"}
	snake := AnimalReview{Food: "mice", Locomotion: "slither", Sound: "hsss"}

	// Handle infinite input
	for {
		fmt.Println("Enter an animal and its information (e.g cow speak): ")
		fmt.Print("> ")
		fmt.Scanf("%s %s", &animal, &info)

		// Check whether the input animal and info are valid
		if !Contains(validAnimal, animal) {
			fmt.Println("Invalid animal. Please enter either cow, bird, or snake")
			fmt.Println()
			continue
		}

		if !Contains(validAnimalInfo, info) {
			fmt.Println("Invalid animal info. Please enter either eat, move, or speak")
			fmt.Println()
			continue
		}

		var chosenAnimal AnimalReview
		switch animal {
		case "cow":
			chosenAnimal = cow
		case "bird":
			chosenAnimal = bird
		case "snake":
			chosenAnimal = snake
		}

		switch info {
		case "eat":
			chosenAnimal.Eat()
		case "move":
			chosenAnimal.Move()
		case "speak":
			chosenAnimal.Speak()
		}

		fmt.Println()
	}
}

func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
