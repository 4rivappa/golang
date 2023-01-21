package main

import "fmt"

func main() {
	fmt.Println("Pointers")

	var inPtr *int
	fmt.Println("pointer initialized with value: ", inPtr)

	number := 23
	pointer := &number
	fmt.Println("value of address: ", pointer)
	fmt.Println("value in address: ", *pointer)

	*pointer += 3
	fmt.Println("number changed to: ", number)
}
