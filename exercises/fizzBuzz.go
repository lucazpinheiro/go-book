// A program that prints the numbers from 1
// to 100. But for multiples of three print "Fizz" in-
// stead of the number and for the multiples of five
// print "Buzz". For numbers which are multiples
// of both three and five print "FizzBuzz".

package main

import (
	"fmt"
)

func main() {
	fizz := "Fizz"
	buzz := "Buzz"
	for i := 1; i <= 100; i++ {
		if isMultipleOfX(i, 3) && isMultipleOfX(i, 5) {
			fmt.Println(fizz + buzz)
			continue
		}
		if isMultipleOfX(i, 3) {
			fmt.Println(fizz)
			continue
		}
		if isMultipleOfX(i, 5) {
			fmt.Println(buzz)
			continue
		}
		fmt.Println(i)
	}
}

func isMultipleOfX(number, x int) bool {
	if number%x == 0 {
		return true
	}
	return false
}
