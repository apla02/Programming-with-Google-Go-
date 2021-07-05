package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a *Animal) Eat() {
	fmt.Println(a.food)
}

func (a *Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a *Animal) Speak() {
	fmt.Println(a.noise)
}

func main() {
	// sliceAnimals := make([]animals, 3)
	cow := Animal{"grass", "walk", "moo"}
	bird := Animal{"worms", "fly", "peep"}
	snake := Animal{"mice", "slither", "hss"}
	// sliceAnimals := [3]Animal{cow, bird, snake}

	fmt.Println("please enter two words: the first one is the name of the animal, either “cow”, “bird”, or “snake”. The second one is the information about the animal, either “Eat”, “Move”, or “Speak”.")
	for {

		fmt.Print(">")
		input := bufio.NewScanner(os.Stdin)
		input.Scan()
		query := strings.Split(input.Text(), " ")
		if len(query) != 2 {
			fmt.Println("Invalid query 1")
		}

		animal, method := query[0], strings.Title(query[1])

		if animal == "cow" {

			defer func() {
				if err := recover(); err != nil {
					fmt.Printf("Please enter a valid action => [%v] not valid\n", method)
				}
			}()
			reflect.ValueOf(&cow).MethodByName(method).Call([]reflect.Value{})

		} else if animal == "bird" {
			defer func() {
				if err := recover(); err != nil {
					fmt.Printf("Please enter a valid action => [%v] not valid\n", method)
				}
			}()
			reflect.ValueOf(&bird).MethodByName(method).Call([]reflect.Value{})
		} else if animal == "snake" {
			defer func() {
				if err := recover(); err != nil {
					fmt.Printf("Please enter a valid action => [%v] not valid\n", method)
				}
			}()
			reflect.ValueOf(&snake).MethodByName(method).Call([]reflect.Value{})
		}

	}
}
