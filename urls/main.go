package main

import (
	"fmt"
	"net/url"
)

const url_one = "https://pkg.go.dev/search?q=net&m=package"

func main() {
	fmt.Println("Urls in GoLang")
	fmt.Println(url_one)

	result, _ := url.Parse(url_one)
	fmt.Println(result.Scheme, result.Host, result.Path)
}
