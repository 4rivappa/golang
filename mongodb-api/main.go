package main

import (
	"fmt"
	"mongodb-api/router"
	"net/http"
)

func main() {
	fmt.Println("MongoDB API - golang")
	r := router.Router()

	// debug
	// fmt.Printf("Router is of type: %T\n", r)
	// fmt.Println(r)

	fmt.Println("Server is getting started...")
	// log.Fatal(http.ListenAndServe(":4000", r))

	err := http.ListenAndServe(":4000", r)
	if err != nil {
		panic(err)
	}
}
