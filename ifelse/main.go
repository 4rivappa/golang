package main

import "fmt"

func main() {
	fmt.Println("If else GoLang")

	loginCount := 23
	var result string

	if loginCount < 10 {
		result = "Regular User"
	} else if loginCount == 10 {
		result = "Exactly 10 login"
	} else {
		result = "another case"
	}

	fmt.Println("Result: ", result)

	// switch case

	// if we want, we can specify fallthrough option
	diceNumber := 4
	switch diceNumber {
	case 1:
		fmt.Println("This is case one")
	case 2:
		fmt.Println("This is case two")
	case 3:
		fmt.Println("This is case three")
	default:
		fmt.Println("This is default option")
	}
}
