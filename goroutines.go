package main

import (
	"fmt"
	"time"
)

// Taken from https://tutorialedge.net/golang/concurrency-with-golang-goroutines/
//What Are Goroutines?
//So to begin with, what are Goroutines? Goroutines are incredibly lightweight “threads” managed by the go runtime.
//They enable us to create asynchronous parallel programs that can execute some tasks far quicker than if they were written in a sequential manner.
//Goroutines are far smaller that threads, they typically take around 2kB of stack space to initialize compared to a thread which takes 1Mb.
//Goroutines are typically multiplexed onto a very small number of OS threads which typically mean concurrent go programs
//require far less resources in order to provide the same level of performance as languages such as Java.
//Creating a thousand goroutines would typically require one or two OS threads at most, whereas if we were to do the same thing in java
//it would require 1,000 full threads each taking a minimum of 1Mb of Heap space.
//By mapping hundreds or thousands of goroutines onto a single thread we don’t have to worry about the performance hit
//when creating and destroying threads in our application. It’s incredibly in-expensive to create and destroy new goroutines
//due to their size and the efficient way that go handles them.

func compute(value int) {
	for i := 0; i < value; i++ {
		time.Sleep(time.Second)
		fmt.Println(i)
	}
}
func main2() {
	// add the ‘go’ keyword in front of our compute function invocation.
	//compute(10)
	//compute(10)
	go compute(10)
	go compute(10)
	select {}
}
