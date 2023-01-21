package main

// using --race parameter to explore the race conditions in our code...!

import (
	"fmt"
	"sync"
)

var numbers = []int{}

func main() {
	fmt.Println("RaceConditions in GoLang")

	wg := &sync.WaitGroup{}

	for i := 0; i < 20; i++ {
		go addNumber(i, wg)
		wg.Add(1)
	}

	wg.Wait()
	fmt.Println(numbers)
}

func addNumber(num int, wg *sync.WaitGroup) {
	defer wg.Done()

	numbers = append(numbers, num)
}

// notes
// Executing without mutex, only using waitgroup
// 	-> There is wait for goroutine completion
// 	-> There is no control of race at the memory level for read and write operations done by GoRoutines

// ## With Mutex we are going to avoid this race conditions for memory read/write
// ## We can also use RWMutex, which we can specifically lock and unlock for read/write operations

// Below is displayed without mutex -> go run --race main.go

// RaceConditions in GoLang
// ==================
// WARNING: DATA RACE
// Read at 0x0000005dd460 by goroutine 8:
//   main.addNumber()
//       C:/Users/ARIVAPPA/personal/explore/go/raceconditions/main.go:27 +0x98
//   main.main.func1()
//       C:/Users/ARIVAPPA/personal/explore/go/raceconditions/main.go:16 +0x47

// Previous write at 0x0000005dd460 by goroutine 7:
//   main.addNumber()
//       C:/Users/ARIVAPPA/personal/explore/go/raceconditions/main.go:27 +0x118
//   main.main.func1()
//       C:/Users/ARIVAPPA/personal/explore/go/raceconditions/main.go:16 +0x47

// Goroutine 8 (running) created at:
//   main.main()
//       C:/Users/ARIVAPPA/personal/explore/go/raceconditions/main.go:16 +0xa6

// Goroutine 7 (finished) created at:
//   main.main()
//       C:/Users/ARIVAPPA/personal/explore/go/raceconditions/main.go:16 +0xa6
// ==================
// ==================
// WARNING: DATA RACE
// Read at 0x00c000100008 by goroutine 9:
//   runtime.growslice()
//       C:/Program Files/Go/src/runtime/slice.go:178 +0x0
//   main.addNumber()
//       C:/Users/ARIVAPPA/personal/explore/go/raceconditions/main.go:27 +0xd2
//   main.main.func1()
//       C:/Users/ARIVAPPA/personal/explore/go/raceconditions/main.go:16 +0x47

// Previous write at 0x00c000100008 by goroutine 7:
//   main.addNumber()
//       C:/Users/ARIVAPPA/personal/explore/go/raceconditions/main.go:27 +0xf3
//   main.main.func1()
//       C:/Users/ARIVAPPA/personal/explore/go/raceconditions/main.go:16 +0x47

// Goroutine 9 (running) created at:
//   main.main()
//       C:/Users/ARIVAPPA/personal/explore/go/raceconditions/main.go:16 +0xa6

// Goroutine 7 (finished) created at:
//   main.main()
//       C:/Users/ARIVAPPA/personal/explore/go/raceconditions/main.go:16 +0xa6
// ==================
// [0 2 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19]
// Found 2 data race(s)
// exit status 66
