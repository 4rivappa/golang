package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("WebReqVerbs Golang")
	// PerformGetRequest()

	// PerformPostRequest()
	PerformPostFormRequest()
}

func PerformPostFormRequest() {
	url_one := "http://localhost:8000/postform"

	data := url.Values{}
	data.Add("firstname", "hitesh")
	data.Add("sort", "alphabet")

	response, err := http.PostForm(url_one, data)
	if err != nil {
		panic(err)
	}
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}

func PerformPostRequest() {
	url := "http://localhost:8000/post"

	requestBody := strings.NewReader(`
	{
		"language": "go",
		"editor": "vscode"
	}
	`)

	response, err := http.Post(url, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}

func PerformGetRequest() {
	url := "http://localhost:8000/get"
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status Code: ", response.StatusCode)
	fmt.Println("Cotent Length: ", response.ContentLength)

	var responseString strings.Builder

	content, _ := ioutil.ReadAll(response.Body)

	byteCount, _ := responseString.Write(content)
	fmt.Println("ByteCount: ", byteCount)
	req_string := responseString.String()

	fmt.Println(req_string)

	// fmt.Println(string(content))
	// another way

}
