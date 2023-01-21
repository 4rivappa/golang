package main

import (
	"fmt"
	"sync"
)

func main() {
	// Refer documentation
	fmt.Println("Channels in GoLang")

	// channels are used to send data messages between the goroutines(parallel processes)

	// make channel which initializes a channel for passing int type objects
	// we can also pass the buffer size of the channel -- refer documentation
	myChannel := make(chan int)

	// left arrow mark usage
	// channel<-val; this is used to send val into channel as it is pointing in to channel
	// var<-channel; this is used to take val out of channel as it is pointing out of channel

	wg := &sync.WaitGroup{}

	wg.Add(2)
	func(ch chan int, wg *sync.WaitGroup) {
		defer wg.Done()

		fmt.Println(<-ch)
	}(myChannel, wg)
	func(ch chan int, wg *sync.WaitGroup) {
		defer wg.Done()

		ch <- 5
		close(ch)
	}(myChannel, wg)

	// there should always be a receiving point for a point to send data
	// closing channel close(channel)

	// Trying Only One Way Channels
	// func(ch <-chan int, wg *sync.WaitGroup) {
	// 	defer wg.Done()

	// 	fmt.Println(<-ch)
	// }(myChannel, wg)
	// func(ch chan<- int, wg *sync.WaitGroup) {
	// 	defer wg.Done()

	// 	ch <- 5
	// 	close(ch)
	// }(myChannel, wg)

	wg.Wait()
}
