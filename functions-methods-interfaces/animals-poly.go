package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Speak()
	Move()
}

type Cow struct {
}

func (cow Cow) Eat() {
	fmt.Println("Grass")
}

func (cow Cow) Move() {
	fmt.Println("Walk")
}

func (cow Cow) Speak() {
	fmt.Println("Moo")
}

type Bird struct {
}

func (bird Bird) Eat() {
	fmt.Println("worms")
}

func (bird Bird) Move() {
	fmt.Println("fly")
}

func (bird Bird) Speak() {
	fmt.Println("peep")
}

type Snake struct {
}

func (snake Snake) Eat() {
	fmt.Println("mice")
}

func (snake Snake) Move() {
	fmt.Println("slither")
}

func (snake Snake) Speak() {
	fmt.Println("hsss")
}

func DoAction(animalAction string, animal Animal) {
	switch animalAction {
	case "move":
		animal.Move()
	case "speak":
		animal.Speak()
	case "eat":
		animal.Eat()
	default:
		fmt.Println("INVALID ACTION")
	}
}

func main() {
	animalsCreated := make(map[string]Animal)
	for {
		var newAnimalType Animal
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">")
		optionsChosen, _ := reader.ReadString('\n')
		optionsText := strings.Split(optionsChosen, " ")
		requestType := strings.ToLower(strings.TrimSpace(optionsText[0]))
		if requestType == "newanimal" {
			animalName := strings.TrimSpace(optionsText[1])
			animalType := strings.TrimSpace(optionsText[2])
			switch animalType {
			case "cow":
				newAnimalType = Cow{}
			case "bird":
				newAnimalType = Bird{}
			case "snake":
				newAnimalType = Snake{}
			default:
				fmt.Println("INVALID ANIMAL")
			}
			animalsCreated[animalName] = newAnimalType
			fmt.Println("Created it!")
		} else if requestType == "query" {
			animalName := strings.TrimSpace(optionsText[1])
			animalAction := strings.TrimSpace(optionsText[2])

			animal, exists := animalsCreated[animalName]
			if !exists {
				fmt.Printf("%s does not exist", animalName)
			} else {
				DoAction(animalAction, animal)
			}

		} else {
			fmt.Println("INVALID REQUEST")
		}
	}

}
