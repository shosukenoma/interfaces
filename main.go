package main

import "fmt"

type bot interface {
	getGreeting() string
}
// if something can do this, then it can be used here
// englishBot can do getGreeting, so englishBot can be used in place of bot

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

// func printGreeting(eb englishBot) {
// 	fmt.Println(eb.getGreeting())
// }

// func printGreeting(sb spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }

// We can omit var name in receiver function if we're not using them in the function.
func (englishBot) getGreeting() string {
	return "Hello!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}