package main

import "fmt"

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

	jim := person{
		firstName: "Jim",
		lastName:  "Alvarez",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 7414,
		},
	}

	jim.updateName("Ima")
	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pointerToPerson *person) updateName(newFristName string) {
	(*pointerToPerson).firstName = newFristName
}
