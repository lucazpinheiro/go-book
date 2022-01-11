package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, r float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (c *Circle) perimetre() float64 {
	return (2 * math.Pi) * c.r
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func (r *Rectangle) perimetre() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return 2 * (l + w)
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

type Shape interface {
	area() float64
	perimetre() float64
}

func printAreas(s ...Shape) {
	for _, shape := range s {
		fmt.Println(fmt.Sprint(shape.area()))
	}
}

func printPerimetres(s ...Shape) {
	for _, shape := range s {
		fmt.Println(fmt.Sprint(shape.perimetre()))
	}
}

func main() {
	c := Circle{15.0, 12.0, 14.0}
	r := Rectangle{5, 5, 1, 1}
	printAreas(&c, &r)
	printPerimetres(&c, &r)
}
