package main

import(
	"fmt"
	"bufio"
	"strings"
	"os"
	"reflect"
	
)

type Animal interface {
	/*This inteface describes the mothods that any kind of animal can do*/
	Eat()
	Move()
	Speak()
}

type Cow struct {
	food, locomotion, noise     string
}
func ( c *Cow) Eat() { fmt.Printf("%v",  c.food)}
func ( c *Cow) Move() {fmt.Printf("%v", c.locomotion)}
func ( c *Cow) Speak() {fmt.Printf("%v", c.noise)}

type Bird struct {
	food, locomotion, noise   string
}
func ( b *Bird) Eat() {fmt.Printf("%v", b.food)}
func ( b *Bird) Move() {fmt.Printf("%v", b.locomotion)}
func ( b *Bird) Speak() {fmt.Printf("%v", b.noise)}

type Snake struct {
	food, locomotion, noise     string
}

func ( s *Snake) Eat() {fmt.Printf("%v",  s.food)}
func ( s *Snake) Move() {fmt.Printf("%v", s.locomotion)}
func ( s *Snake) Speak() {fmt.Printf("%v", s.noise)}


func main() {
	// sliceAnimals := make([]animals, 3)
	var nameType map[string]string

	nameType = make(map[string]string)
	// sliceAnimals := [3]Animal{cow, bird, snake}

	fmt.Println("command from the user must be either a “newanimal” command or a “query” command")
	for {

		fmt.Print(">")
		fmt.Println("please enter name of command ( 'newanimal' or 'query), a name of animal (cow, bird or snake) and one information (eat, move or speak) separted by an blank space")
		
		input := bufio.NewScanner(os.Stdin)
		input.Scan()
		query := strings.Split(input.Text(), " ")
		if len(query) != 3 {
			fmt.Println("Invalid input, must be 3 words")
		}

		command := query[0]

		switch command {
		case "newanimal":

			name, animalType := query[1], query[2]

			fmt.Printf("name: %s\n", name)
			fmt.Printf("animalType: %s\n", animalType)

			//adding name checking
			_, ok := nameType[name]
			if ok == false {
				nameType[name] = animalType
				fmt.Println("Animal was created")
			} else {
				fmt.Printf("Please try again", name)
			}
		case "query":

			
			cow := Cow{"grass", "walk", "moo"}
			bird := Bird{"worms", "fly", "peep"}
			snake := Snake{"mice", "slither", "hss"}
			
			animal, method := query[1], strings.Title(query[2])
			
			fmt.Printf("name of animal: %s\n", animal)
			fmt.Printf("info requerided: %s\n", method)
			if animal == "cow" {

				defer func() {
					if err := recover(); err != nil {
						fmt.Printf("Please enter a valid action => [%v] not valid\n", method)
					}
				}()
				reflect.ValueOf(&cow).MethodByName(method).Call([]reflect.Value{})
				fmt.Println("")

			} else if animal == "bird" {
				defer func() {
					if err := recover(); err != nil {
						fmt.Printf("Please enter a valid action => [%v] not valid\n", method)
					}
				}()
				reflect.ValueOf(&bird).MethodByName(method).Call([]reflect.Value{})
				fmt.Println("")
			} else if animal == "snake" {
				defer func() {
					if err := recover(); err != nil {
						fmt.Printf("Please enter a valid action => [%v] not valid\n", method)
					}
				}()
				reflect.ValueOf(&snake).MethodByName(method).Call([]reflect.Value{})
				fmt.Println("")
			}
	}
}
}