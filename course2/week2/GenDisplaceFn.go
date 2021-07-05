/*
Write a program which first prompts the user
to enter values for acceleration, initial velocity, and initial displacement.
Then the program should prompt the user to enter a value for time and the
program should compute the displacement after the entered time.

You will need to define and use a function called GenDisplaceFn() which takes three float64
arguments, acceleration a, initial velocity vo, and initial displacement so. GenDisplaceFn()
should return a function which computes displacement as a function of time, assuming the given
values acceleration, initial velocity, and initial displacement. The function returned by GenDisplaceFn()
 should take one float64 argument t, representing time, and return one
float64 argument which is the displacement travelled after time t.
*/
package main

import (
	"fmt"
	"math"
)

// s = ½ a t2 + vot + so
func GenDisplaceFn(a, vo, so float64) func(float64) float64 {

	DisplaceFn := func(t float64) float64 {
		t2 := math.Pow(t, 2)
		s := (1 / 2 * a * t2) + (vo * t) + so
		return s
	}
	return DisplaceFn

}

func main() {
	var values []float64
	fmt.Println("Please enter 3 floats separated by a blank space =>>  a: acceleration, so: initial displacement and vo: inicial velocity")
	for i := 0; i < 3; i++ {
		var number float64
		_, err := fmt.Scan(&number)
		if err != nil {
			fmt.Println("Please enter a valid number")
		}
		values = append(values, number)

	}
	fmt.Printf("The values entered are [a]: %v, [so]:%v and [vo]: %v\n", values[0], values[1], values[2])

	fn := GenDisplaceFn(values[0], values[1], values[2])

	fmt.Println("Please enter a time to calcule the displacement")
	var t float64
	_, err := fmt.Scan(&t)
	if err != nil {
		fmt.Println("Please enter a valid number")
	}
	fmt.Printf("the displacement after %v seconds is %v\n", t, fn(t))
}
