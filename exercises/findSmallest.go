// finds the smallest number in a given list

package main

import "fmt"

var list = []int{
	48, 96, 86, 68,
	57, 82, 63, 70,
	37, 34, 83, 27,
	19, 97, 9, 17,
}

func main() {
	position, value := findSmallest(list)
	fmt.Println("Smallest value: ", +value, ", position: ", position)
}

func findSmallest(list []int) (int, int) {
	position := 0
	smallest := list[position]
	for i, number := range list {
		if number < smallest {
			position = i
			smallest = number
		}
	}
	return position, smallest
}
