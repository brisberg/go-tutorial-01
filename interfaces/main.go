package main

import "fmt"

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

func (englishBot) getGreeting() string {
	return "Hi there!" // custom code for englishbot
}

func (spanishBot) getGreeting() string {
	return "Hola!" // custom code for spanishbot
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
