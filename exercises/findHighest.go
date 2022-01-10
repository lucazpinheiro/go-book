// Write a function with one variadic parameter
// that finds the greatest number in a list of num-
// bers.

package main

import (
	"fmt"
)

/*
... variadiac parameter, indicates a function can take zero or many arguments.
*/

func findHighest(args ...int) int {
	highest := args[0]
	for _, number := range args {
		if highest < number {
			highest = number
		}
	}
	return highest
}

func main() {
	fmt.Println(findHighest(40, 1, 8, 98))
	fmt.Println(findHighest(4, 9, 34, 32, 11))
}
