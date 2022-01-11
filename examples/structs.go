package main

import "fmt"

type Address struct {
	street      string
	houseNumber int
}

func (a *Address) FullAddress() string {
	return a.street + ", " + fmt.Sprint(a.houseNumber) + "."
}

type Person struct {
	fname     string
	lname     string
	birthYear int
	Address
}

func (p *Person) FullName() string {
	return p.fname + " " + p.lname
}

func (p *Person) Age() int {
	return 2022 - p.birthYear
}

func main() {
	someone := Person{"lucas", "pinheiro", 1999, Address{"Rua qualquer", 100}}
	fmt.Println(someone.FullName(), "is", someone.Age(), "years old, and leave on", someone.FullAddress())
}

// using struct types passing them to "alone" function/procedures
// type Person struct {
// 	fname     string
// 	lname     string
// 	birthYear int
// 	age       int
// }

// func getFullName(p Person) {
// 	fmt.Println(p.fname + " " + p.lname)
// }

// func calculateAge(p *Person) {
// 	p.age = 2022 - p.birthYear
// 	fmt.Println(p.age)
// }

// func main() {
// 	someone := Person{"lucas", "pinheiro", 1999, 0}
// 	// getFullName(someone)
// 	fmt.Println(someone.fname, " is ", someone.age, " years old")
// 	calculateAge(&someone)
// 	fmt.Println(someone.fname, " is ", someone.age, " years old")
// }
