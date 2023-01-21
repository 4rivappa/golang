package main

import "fmt"

func main() {
	fmt.Println("Maps in GoLang")

	languages := make(map[string]string)
	languages["js"] = "javascript"
	languages["py"] = "python"
	languages["rs"] = "rust"

	fmt.Println("Languages: ", languages)
	fmt.Println("py stands for: ", languages["py"])

	delete(languages, "py")
	fmt.Println("Languages: ", languages)

	// loops for map
	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
}
