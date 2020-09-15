package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var urls = []string{
	"https://google.com",
	"https://tutorialedge.net",
	"https://twitter.com",
}

func fetch(url string, wg *sync.WaitGroup) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	wg.Done() // mark goroutine as completed
	fmt.Println(resp.Status)
	return resp.Status, nil
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HomePage Endpoint Hit")
	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1) // ADD 1 every time adding a goroutine
		go fetch(url, &wg)
	}

	wg.Wait() // wait for the group goroutines for completition
	fmt.Println("Returning Response")
	fmt.Fprintf(w, "Responses")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequests()
}
