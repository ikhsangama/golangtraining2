package main

import "fmt"

type bot interface {
	// If you are a type in this program with a
	// function called 'getGreeting' and you return a
	// string then you are now an honorary member
	// of type 'bot'
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
