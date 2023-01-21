package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("Json - Handling")

	// EncodeJson()
	DecodeJson()
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	    {
			"coursename": "Rust",
            "price": 20,
            "website": "Udemy",
            "tags": [
                "Backend",
                "cargo"
            ]
		}
	`)

	var lcoCourse course
	if json.Valid(jsonDataFromWeb) {
		fmt.Println("Json was valid..")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("Json was not valid...")
	}

	// storing it in a map
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)
}

func EncodeJson() {
	lcoCourses := []course{
		{"Reactjs", 0, "Udemy", "HereWeGo", []string{"Web", "Facebook"}},
		{"VewJs", 10, "Udemy", "HereWeGo", nil},
		{"Rust", 20, "Udemy", "HereWeGo", []string{"Backend", "cargo"}},
	}

	// finalJson, err := json.Marshal(lcoCourses)
	finalJson, err := json.MarshalIndent(lcoCourses, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(finalJson))
}

type course struct {
	Name     string   `json:"coursename"`
	Price    int      `json:"price"`
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
	// `` is setting alias for the given field name
	// "-", will remove the entire element, while converting
	// ",omitempty", will remove the element if it is nil
}
