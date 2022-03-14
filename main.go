package main

import "fmt"

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
	jim := person{
		firstName: "Jim",
		lastName:  "Parse",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 1323,
		},
	}
	fmt.Printf("%+v", jim)
}
