package main

import "fmt"

// NOTE: interface used to allow this functionality with all types that match the interface
// aka the types that have a getGreeting() method that has a receiver and returns a string
// * this will not work if there is a different return type for getGreeting()
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

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// NOTE: if not using variable in receiver, you can delte variable and just use the required type
// originally: (eb englishBot) --> (englishBot)
func (englishBot) getGreeting() string {
	// custom logic that can't be shared
	return "Hi there"
}

func (spanishBot) getGreeting() string {
	// custom logic that can't be shared
	return "Hola"
}
