package main

import "fmt"

func main() {
	// there is value types : int float string bool structs
	// use pointers to change things in func

	// and there is reference type ( still referencing the same underlying source of data in memory) : slices maps channels pointers functions
	// don't worry about pointers

	mySlice := []string{"Hi", "There"}
	fmt.Printf("pointer 1 %p \n", &mySlice)
	fmt.Printf("pointer 1 %p \n", &mySlice[0])
	updateSlice(mySlice)
	fmt.Println(mySlice)
}

func updateSlice(s []string) {
	fmt.Printf("pointer 2 %p \n", &s)
	fmt.Printf("pointer 2 %p \n", &s[0])
	s[0] = "Bye"
}
