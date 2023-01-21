package main

import (
	"fmt"
	"net/http"
	"sync"
)

// keyword 'go' for creating goroutine

var signals = []string{"test"}

var wg sync.WaitGroup
var mutex sync.Mutex

func main() {
	fmt.Println("GoRoutines in GoLang")

	// go greeter("hello")
	// greeter("world")
	// In the above two lines code, we are not waiting for the goroutine to return or complete task...

	websites := []string{
		"http://lco.dev",
		"http://go.dev",
		"http://you.com",
		"http://bing.com",
		"http://google.com",
	}

	for _, website := range websites {
		go getStatusCode(website)
		wg.Add(1)
	}

	wg.Wait()
	fmt.Println(signals)
}

func getStatusCode(url string) {
	defer wg.Done()

	result, err := http.Get(url)
	if err != nil {
		fmt.Println("Error in getting ", url)
	} else {
		mutex.Lock()
		signals = append(signals, url)
		mutex.Unlock()

		fmt.Printf("Website %s got status code %d\n", url, result.StatusCode)
	}
}

func greeter(s string) {
	for i := 0; i < 6; i++ {
		fmt.Println(s)
	}
}
