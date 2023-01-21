package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Files in GoLang")

	content := "Here we go, now it's time for file read/write"

	file, err := os.Create("./test.txt")
	checkNilError(err)

	length, err := io.WriteString(file, content)
	checkNilError(err)

	fmt.Println("length is: ", length)

	readFile("./test.txt")
	defer file.Close()
}

func readFile(filepath string) {
	databyte, err := ioutil.ReadFile(filepath)
	checkNilError(err)

	fmt.Println("The output string is \n", string(databyte))
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
