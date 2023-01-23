package main

import (
	"fmt"
	"os"
	"strconv"
)

// https://gobyexample.com/command-line-arguments
func main() {

	var secretNumber int = 3

	guessedNumber, _ := strconv.Atoi(os.Args[1])

	if secretNumber == guessedNumber {
		fmt.Println("You discovered the secret number")
	} else {
		fmt.Println("Too bad try again")
	}
}
