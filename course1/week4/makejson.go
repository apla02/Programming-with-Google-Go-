/*
Write a program which prompts the user to first enter a name, and then enter an address.
Your program should create a map and add the name and address to the map using the keys
“name” and “address”, respectively. Your program should use Marshal() to create a JSON object
from the map, and then your program should print the JSON object.
*/

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("=> Please give your name and then press enter")
	name, _ := reader.ReadString('\n')
	fmt.Println("=> Please give your address and then press enter")
	address, _ := reader.ReadString('\n')

	if len(name) > 0 && len(address) > 0 {
		objGo := make(map[string]string)
		objGo["name"] = name[:len(name)-1]
		objGo["address"] = address[:len(name)-1]
		obj, err := json.Marshal(objGo)
		if err != nil {
			log.Fatal("error in json to byte serialization. err is ", err)
			return
		}
		fmt.Println(string(obj))
	}
}
