package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Reads a file called triangles.txt and tells which lines represents a triangle
func main() {

	textFile, err := os.Open("triangles.txt")

	if err != nil {
		panic(err)
	}

	textFileScanner := bufio.NewScanner(textFile)
	textFileScanner.Split(bufio.ScanLines)

	for textFileScanner.Scan() {
		var line = textFileScanner.Text()
		// https://www.geeksforgeeks.org/how-to-split-a-string-in-golang/
		var strNumbers = strings.Split(line, " ")
		if len(strNumbers) != 3 {
			fmt.Printf("line [%s] isn't valid and will be ignored\n", line)
		} else {
			a, _ := strconv.Atoi(strNumbers[0])
			b, _ := strconv.Atoi(strNumbers[1])
			c, _ := strconv.Atoi(strNumbers[2])
			if a+b > c && a+c > b && b+c > a {
				fmt.Printf("%s can form a triangle\n", strNumbers)
			} else {
				fmt.Printf("%s can't form a triangle\n", strNumbers)
			}
		}
	}

	_ = textFile.Close()
}
