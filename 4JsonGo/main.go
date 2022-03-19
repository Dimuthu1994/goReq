package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int `json:"price"`
	Platform string `json:"website"`
	Password string `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	EncodeJson()
}

func EncodeJson() {
	myCourses := []course{
		{"React", 299, "youtube", "abc123", []string{"web-dev", "js"}},
		{"Node", 199, "youtube", "asdfs3", []string{"web-dev", "js"}},
		{"JS", 399, "youtube", "fdhgdth23", nil},
	}

	//package this data as JSON data

	//Marshall-implement JSON
	// we get null when empty (nil)
	//finalJson,err:= json.Marshal(myCourses)

	finalJson,err:= json.MarshalIndent(myCourses,"","\t")

	if err!=nil{
		panic(err)
	}
	fmt.Printf("%s\n",finalJson)
}