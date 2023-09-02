package main

import (
	"fmt"
	"time"
)

var (
	balance            float64 = 0
	transactionChannel         = make(chan float64)
)

func doTransaction(amount float64) {
	balance += amount
}

func doChanTransaction() {
	balance += <-transactionChannel
}

func main() {

	for i := 1; i < 1000; i++ {
		go doTransaction(100)
	}

	for i := 1; i < 1000; i++ {
		go doTransaction(-100)
	}

	time.Sleep(time.Duration(3) * time.Second)
	fmt.Printf("Resulting balance without channels: %f\n", balance)

	balance = 0

	for i := 1; i < 1000; i++ {
		go doChanTransaction()
		transactionChannel <- 100
	}

	for i := 1; i < 1000; i++ {
		go doChanTransaction()
		transactionChannel <- -100
	}

	time.Sleep(time.Duration(3) * time.Second)
	fmt.Printf("Resulting balance with channels: %f\n", balance)

	// not mandatory here because the program is about to finish anyway
	// but remember to close your channels to signal the work is done
	close(transactionChannel)
}
