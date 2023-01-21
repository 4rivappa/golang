package main

import "fmt"

func main() {
	fmt.Println("Arrays in Golang")

	var fruitList [4]string
	fruitList[1] = "apple"

	fmt.Println("Here is the array: ", fruitList)
	fmt.Println("Length of array: ", len(fruitList))

	var vegList = [3]string{"Potato", "Beans", "Carrot"}
	fmt.Println("Veg List is: ", vegList)
	fmt.Println("Length of veg list: ", len(vegList))
}
