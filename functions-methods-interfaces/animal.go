package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (animal *Animal) Eat() string {
	return animal.food
}

func (animal *Animal) Move() string {
	return animal.locomotion
}

func (animal *Animal) Speak() string {
	return animal.noise
}

func main() {

	for {
		var animal Animal

		fmt.Println("Please insert an animal with an action separated by an space: ")
		fmt.Println("Possible animals:\n1. Cow\n2. Bird\n3. Snake")
		fmt.Println("Possible actions:\n1. Eat\n2. Move\n3. Speak")

		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">")
		optionsChosen, _ := reader.ReadString('\n')
		optionsText := strings.Split(optionsChosen, " ")

		animalChosen := strings.ToLower(strings.TrimSpace(optionsText[0]))
		functionChosen := strings.ToLower(strings.TrimSpace(optionsText[1]))

		switch animalChosen {
		case "cow":
			animal = Animal{"grass", "walk", "moo"}
		case "bird":
			animal = Animal{"worms", "fly", "peep"}
		case "snake":
			animal = Animal{"mice", "slither", "hsss"}
		}

		switch functionChosen {
		case "eat":
			fmt.Println("Animal Food: ", animal.Eat())
		case "move":
			fmt.Println("Animal locomotion: ", animal.Move())
		case "speak":
			fmt.Println("Animal noise: ", animal.Speak())
		}

	}
}
