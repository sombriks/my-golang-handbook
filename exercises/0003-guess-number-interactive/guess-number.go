package main

import "fmt"

// https://pkg.go.dev/fmt#Scanln
func main() {

	var secretNumber int = 3

	var guessedNumber int

	fmt.Println("Guess a number")
	fmt.Scanln(&guessedNumber)

	if secretNumber == guessedNumber {
		fmt.Println("You discovered the secret number")
	} else {
		fmt.Println("Too bad try again")
	}
}
