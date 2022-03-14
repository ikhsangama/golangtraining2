package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	firstWay := person{"ikhsan", "gama"}
	fmt.Println(firstWay)

	secondWay := person{firstName: "ikhsan", lastName: "gama"}
	fmt.Println(secondWay)

	var thirdWay person
	fmt.Println(thirdWay)
	thirdWay.firstName = "ikhsan"
	thirdWay.lastName = "gama"
	fmt.Printf("%+v", thirdWay)
}
