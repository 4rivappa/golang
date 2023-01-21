package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slices Golang")

	var fruitList = []string{"Apple", "Peach", "Watermelon"}
	fmt.Printf("fruiList type is: %T\n", fruitList)
	fmt.Println("Fruits list: ", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println("Fruits list after apppend: ", fruitList)

	//slice the slices type
	fruitList = append(fruitList[:3])
	fmt.Println("Fruits list after :3 slice ", fruitList)

	highScores := make([]int, 4)
	highScores[1] = 234
	highScores[0] = 123
	highScores[2] = 938

	// highScores[4] = 233 // error index out of bound
	fmt.Println("High scores list: ", highScores)
	highScores = append(highScores, 444, 555)
	fmt.Println("New Updated scores: ", highScores)

	sort.Ints(highScores)
	fmt.Println("Sorted highscores: ", highScores)

	// how to remove a value from slices based on index
	var courses = []string{"reactjs", "python", "ruby", "javascript"}
	fmt.Println("Courses slices: ", courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println("Removed courses slice: ", courses)

}
