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
		lastName:  "Parse",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 1323,
		},
	}

	jim.updateName("jimmy")
	jim.print()
}

func (p person) updateName(newFirstname string) {
	p.firstName = newFirstname
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
