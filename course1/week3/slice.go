/*
 Write a program which prompts the user to enter integers and stores the integers in a sorted slice.
 The program should be written as a loop. Before entering the loop, the program should create an empty
 integer slice of size (length) 3. During each pass through the loop, the program prompts the user to enter
 an integer to be added to the slice. The program adds the integer to the slice, sorts the slice, and prints
 the contents of the slice in sorted order. The slice must grow in size to accommodate any number of integers
 which the user decides to enter. The program should only quit (exiting the loop) when the user enters the character
 ‘X’ instead of an integer.
*/
package main

import (
	"fmt"
	"sort"
	"strconv"
)

const (
	size       int    = 3
	exitPrompt string = "X"
)

func main() {
	sliceInt := make([]int, size)
	var valInput string

	for {
		fmt.Println("=> Please enter a number to be added to a slice or 'X' to exit")
		fmt.Scan(&valInput)
		if valInput == exitPrompt {
			break
		}
		sliceVal, err := strconv.Atoi(valInput)
		if err != nil {
			fmt.Println("Please enter a valid number")
			continue
		}
		sliceInt = append(sliceInt, sliceVal)
		sort.Ints(sliceInt)
		var result []int = sliceInt[size:len(sliceInt)]
		// fmt.Println(sliceInt) print with empty slots
		fmt.Println(result)
	}

}
