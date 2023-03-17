package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"time"
)

func _print(delay int, count int) {

	for count != 0 {
		count--
		fmt.Println(count)
		time.Sleep(time.Duration(delay) * time.Second)
	}

	fmt.Println("_print done!")
}

// exmaple go run concurrent-printers.go 1 2 10
func main() {

	if len(os.Args) < 4 {
		fmt.Printf("Usage: %s <delay 1> <delay 2> <count>\n", os.Args[0])
		return
	}

	d1, _ := strconv.Atoi(os.Args[1])
	d2, _ := strconv.Atoi(os.Args[2])
	count, _ := strconv.Atoi(os.Args[3])

	go _print(d1, count)
	go _print(d2, count)

	// needed so main thread don't finish before the two go routines
	time.Sleep(time.Second * time.Duration((count+1)*int(math.Max(float64(d1), float64(d2)))))
}
