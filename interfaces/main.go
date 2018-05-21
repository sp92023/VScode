package main

import (
	"fmt"
)

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

/*func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

func printGreeting(sb spanishBot) {
	fmt.Println(sb.getGreeting())
}*/
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// func (eb englishBot)getGreeting() string {}
func (englishBot) getGreeting() string {
	// very custom logic for generating an english greeting
	return "Hi There!"
}

// func (spanishBot)getGreeting() string {}
func (sb spanishBot) getGreeting() string {
	return "Hola!"
}
