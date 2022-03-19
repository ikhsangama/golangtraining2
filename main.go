package main

import "fmt"

type englishBot struct {
}
type spanishBot struct {
}

func main() {
	eb := englishBot{}
	//sb := spanishBot{}

	printGreeting(eb)
	//printGreeting(sb)
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

func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

//func printGreeting(sb spanishBot) {
//	fmt.Println(sb.getGreeting())
//}
