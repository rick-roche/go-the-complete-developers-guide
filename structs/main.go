package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	rick := person{
		firstName: "Rick",
		lastName:  "Roche",
		contactInfo: contactInfo{
			email:   "test@test.com",
			zipCode: 12345,
		},
	}

	// rickPointer := &rick
	// rickPointer.updateName("Richard")
	rick.updateName("Richard")
	rick.print()

	// value types: [use pointers to change these in functions]
	// - int, float, string, bool
	// - structs

	// referenced types: [dont worry about pointers with these]
	// - slices
	// - maps
	// - channels
	// - pointers
	// - functions

}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
