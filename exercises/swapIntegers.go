// Write a program that can swap two integers
// ( x := 1; y := 2; swap(&x, &y) should give you
// x=2 and y=1 ).

package main

import "fmt"

func swap(x, y *int) {
	*y, *x = *x, *y
}

func main() {
	x, y := 10, 333
	fmt.Println("x: ", x, " y: ", y)
	fmt.Println("swapping...")
	swap(&x, &y)
	fmt.Println("x: ", x, " y: ", y)
}
