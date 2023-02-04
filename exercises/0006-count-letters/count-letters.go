package main

import (
	"fmt"
	"os"
)

func main() {

	word := os.Args[1]
	letters := map[string]int{}
	fmt.Println(letters, " capacity: ", len(letters))

	for _, c := range word {
		l := string(c)
		_, present := letters[l]
		if present {
			letters[l] += 1
		} else {
			letters[l] = 1
		}
	}

	fmt.Println(letters, " capacity: ", len(letters))

}
