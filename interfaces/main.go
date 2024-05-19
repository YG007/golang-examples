package main

import "fmt"

// first example to initilize an interface
type bot0 interface {
	getGreeting() string
}

// second example to initilize an interface
type user struct {
	name string
}

type bot1 interface {
	getGreeting(string, int) (string, error)
	getBotVersion() float64
	respondToUse(user) string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot0) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	// very custom logic for generating an english greeting
	return "Hi There!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}
