package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Time")

	currTime := time.Now()
	fmt.Println(currTime)

	fmt.Println(currTime.Format("01-02-2006 Monday"))
}
