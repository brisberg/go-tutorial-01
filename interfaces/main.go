package main

type englishBot struct{}
type spanishBot struct{}

func main() {

}

func (englishBot) getGreeting() string {
	return "Hi there!" // custom code for englishbot
}

func (spanishBot) getGreeting() string {
	return "Hola!" // custom code for spanishbot
}
