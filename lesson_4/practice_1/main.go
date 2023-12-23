package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	newBot := englishBot{}
	fmt.Println(newBot.getGreeting())
	printGreeting(newBot)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	// Imagine we wrote VERY different logic to implement out data
	return "Hi there"
}

func (spanishBot) getGreeting() string {
	return ""
}
