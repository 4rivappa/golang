package main

import "fmt"

func main() {
	fmt.Println("Structs in GoLang")

	// no inheritance in golang;
	// no super, or parent kind of stuff;

	hitesh := User{"Hitesh", 20, true, "hitesh@go.dev"}
	fmt.Println(hitesh)
	fmt.Printf("hitesh details are: %+v \n", hitesh)
	fmt.Printf("name is %v", hitesh.Name)

	hitesh.GetStatus()
}

type User struct {
	Name   string
	Age    int
	Status bool
	Email  string
}

// methods in golang

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}
