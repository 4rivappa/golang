package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("User Input")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your rating for the go:")

	// comma, ok || comma, err -> sytax
	input, _ := reader.ReadString('\n')
	fmt.Println(input)
	fmt.Printf("Type of the input is: %T", input)
}
