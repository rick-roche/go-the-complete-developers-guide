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
	contact   contactInfo
}

func main() {
	rick := person{
		firstName: "Rick",
		lastName:  "Roche",
		contact: contactInfo{
			email:   "test@test.com",
			zipCode: 12345,
		},
	}

	// var rick person
	// rick.firstName = "Rick"
	// rick.lastName = "Roche"
	// rick.contact.email = "test@test.com"
	// rick.contact.zipCode = 1234

	// fmt.Println(rick)
	fmt.Printf("%+v", rick)
}
