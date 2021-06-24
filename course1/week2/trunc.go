package main

import(
	"fmt"
	"log"
)

func main(){
/*
Write a program which prompts the user to enter a floating point number and prints the integer 
which is a truncated version of the floating point number that was entered.
Truncation is the process of removing the digits to the right of the decimal place.
*/
var floatNumber float64
fmt.Println("Please enter a float number:=> ")
_, err := fmt.Scan(&floatNumber)
if err != nil {
	log.Printf("[x] Value entered is invalid !")
}
fmt.Printf("The number entered  was '%v' and converted in a integer is %v\n", floatNumber, int64(floatNumber))
}