package main

import "fmt"

func main() {
	numbers := make([]int, 11)
	for index, _ := range numbers {
		if index%2 == 0 {
			fmt.Printf("%v is even \n", index)
		} else {
			fmt.Printf("%v is odd \n", index)
		}
	}

}
