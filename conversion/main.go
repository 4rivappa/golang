package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Conversion")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your rating for the go:")

	// comma, ok || comma, err -> sytax
	input, _ := reader.ReadString('\n')
	fmt.Println(input)
	fmt.Printf("Type of the input is: %T \n", input)

	// numRating, err := strconv.ParseFloat(input, 64)
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		numRating += 1
		fmt.Println("Added one to your rating: ", numRating)
	}
}
