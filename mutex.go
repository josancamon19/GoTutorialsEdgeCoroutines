package main

import (
	"fmt"
	"sync"
)

var (
	mutex   sync.Mutex
	balance int
)

func init() {
	balance = 100
}

// Explanation here https://tutorialedge.net/golang/go-mutex-tutorial/
// 1. ensure that regardless of how your goroutine terminates
// 2.  Unlock() on an error, then it is possible that your application will go into a deadlock as other goroutines will be unable to attain
// Check the code explanation

func deposit(value int, wg *sync.WaitGroup) {
	mutex.Lock()
	fmt.Printf("Depositing %d to account with balance: %d\n", value, balance)
	balance += value
	mutex.Unlock()
	wg.Done()
}

func withdraw(value int, wg *sync.WaitGroup) {
	mutex.Lock()
	fmt.Printf("Withdrawing %d from account with balance: %d\n", value, balance)
	balance -= value
	mutex.Unlock()
	wg.Done()
}
func main3() {
	fmt.Println("Go Mutex Example")

	var wg sync.WaitGroup
	wg.Add(2)
	go withdraw(700, &wg)
	go deposit(500, &wg)
	wg.Wait()

	fmt.Printf("New Balance %d\n", balance)
}
