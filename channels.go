package main

import (
	"fmt"
	"math/rand"
	"time"
)

// myChannel := make(chan int) - creates myChannel which is a channel of type int
// channel <- value - sends a value to a channel
// value := <- channel - receives a value from a channel

func CalculateValue(values chan int) {
	value := rand.Intn(10)
	fmt.Println("Calculated Random Value: {}", value)
	time.Sleep(1 * time.Second)
	values <- value
}

func bufferedChannel() {
	fmt.Println("Go Channel Tutorial")

	values := make(chan int)
	defer close(values) // close channel when main function done

	go CalculateValue(values)
	go CalculateValue(values)

	value := <-values
	fmt.Println(value)
	time.Sleep(1 * time.Second)

}

func unBufferedChannel() {
	fmt.Println("Go Channel Tutorial")

	values := make(chan int, 2)
	defer close(values) // close channel when main function done

	go CalculateValue(values)
	go CalculateValue(values)

	value := <-values
	fmt.Println(value)
	time.Sleep(1 * time.Second)
}

func main4() {
	unBufferedChannel()
}
