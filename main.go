package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipCode int
}

//struct inside a struct
type person struct {
	firstName string
	lastName  string
	//contact   contactInfo
	contactInfo
}

func main() {
	/*alex := person{firstName: "Alex", lastName: "Anderson"}
	fmt.Println(alex)

	var alex person

	alex.firstName = "Alex"
	alex.lastName = "Anderson"

	fmt.Println(alex)
	fmt.Printf("%+v", alex) */

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	jim.print()

}

func (p person) print() {
	fmt.Printf("%+v", p)
}
