package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	balance float64     = 0
	mutex   sync.Locker = &sync.Mutex{}
)

func doTransaction(amount float64) {
	balance += amount
}

func doMutexTransaction(amount float64) {
	mutex.Lock()
	balance += amount
	mutex.Unlock()
}

func main() {

	for i := 1; i < 1000; i++ {
		go doTransaction(100)
	}

	for i := 1; i < 1000; i++ {
		go doTransaction(-100)
	}

	time.Sleep(time.Duration(3) * time.Second)
	fmt.Printf("Resulting balance without mutext: %f\n", balance)

	balance = 0

	for i := 1; i < 1000; i++ {
		go doMutexTransaction(100)
	}

	for i := 1; i < 1000; i++ {
		go doMutexTransaction(-100)
	}

	time.Sleep(time.Duration(3) * time.Second)
	fmt.Printf("Resulting balance with mutext: %f\n", balance)
}
