/*Review criteria
Write a program which prompts the user to enter a string.
The program searches through the entered string for the characters ‘i’, ‘a’, and ‘n’.
The program should print “Found!” if the entered string starts with the character ‘i’,
ends with the character ‘n’, and contains the character ‘a’. The program should print “Not Found!” otherwise.
The program should not be case-sensitive, so it does not matter if the characters are upper-case or lower-case.

Examples: The program should print “Found!” for the following example entered strings, “ian”, “Ian”, “iuiygaygn”, “I d skd a efju N”.
The program should print “Not Found!” for the following strings, “ihhhhhn”, “ina”,
 “xian”ll be given for the second test execution, if the program correctly prints "Not Found!"
when a string that does not contain the characters ‘i’, ‘a’, and ‘n’ is entered.
*/

package main

import (
	"fmt"
	"strings"
	"log"
)

func main(){
	var text string
 	fmt.Println("Please enter a string that contains characteris 'i', 'a', or 'n' :")
  	fmt.Println("---------------------")

	_, err := fmt.Scan(&text)
	if err != nil {
		log.Printf("[x] Value entered is invalid !")
	}

	count:= 0
	if text[0] == 'I' || text[0] == 'i'{
		count += 1
	}
	if strings.ContainsAny(text, "aA"){
		count += 1
	}
	if text[len(text)-1] == 'n' || text[len(text)-1]  == 'N'{
		count += 1
	}
	if count == 3{fmt.Println("Found!")} else{
		fmt.Println("Not Found!")

	}


}