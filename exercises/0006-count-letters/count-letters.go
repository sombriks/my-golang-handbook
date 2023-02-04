package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	word := os.Args[1]
	letters := map[string]int{}
	fmt.Println(letters, " capacity: ", len(letters))
	list := strings.Split(word, "")

	for i := range list {
		_, present := letters[list[i]]
		if present {
			letters[list[i]] += 1
		} else {
			letters[list[i]] = 1
		}
	}

	fmt.Println(letters, " capacity: ", len(letters))

}
