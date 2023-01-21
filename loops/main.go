package main

import "fmt"

func main() {
	fmt.Println("Loops in GoLang")

	days := []string{"Monday", "Tuesday", "Wednesday", "Saturday"}

	// type one
	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	// type two
	for i := range days {
		fmt.Println(days[i])
	}

	// type three
	for index, day := range days {
		fmt.Printf("The index %v, value is %v \n", index, day)
	}

	rougeValue := 1
	for rougeValue < 10 {
		if rougeValue == 2 {
			goto lco
		}

		if rougeValue == 5 {
			break
		}
		fmt.Println("Value ", rougeValue)
		rougeValue += 1

		if rougeValue < 1 {
			continue
		}
	}

lco:
	fmt.Println("This is the goto place for LCO")
}
