package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct {
}
type spanishBot struct {
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// you can omit the value of struct if not using it
func (englishBot) getGreeting() string {
	//	CUSTOM LOGIC
	return "Hi There"
}

// you can omit the value of struct if not using it
func (spanishBot) getGreeting() string {
	//	CUSTOM LOGIC
	return "Hola"
}
