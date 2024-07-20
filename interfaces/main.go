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

func printGreeting(eb englishBot) { //this line does not contain a return value because it will only take input, then print something out.
	fmt.Println(eb.getGreeting())
}

func printGreeting(sb spanishBot) {
	fmt.Println(sb.getGreeting())
}

func (eb englishBot) getGreeting() string {
	//Receiver Value:  eb, of type englishBot
	//function name:  getGreeting
	//Returns:  a string
	//VERY cutom logic for generating an english greeting
	return "Hi There!"

}

func (sb spanishBot) getGreeting() string {
	return "Hola!"
}
