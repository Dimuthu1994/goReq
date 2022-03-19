package main

import (
	"fmt"
	"net/url"
)
const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=dsffgsfg"

func main() {
	fmt.Println("Handling URLS") 
	fmt.Println(myurl)

	//parsing
	result,_ := url.Parse(myurl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)


	qparams := result.Query()
	//key values 
	fmt.Printf("qparams are : %T\n",qparams)
	fmt.Println(qparams["coursename"])

	for _,val := range qparams{
		fmt.Println("params is:",val)
	}

	//construct a url
	partsOfUrl := &url.URL{
		Scheme: "https",
		Host:"lco.dev",
		Path:"/tutcss",
		RawPath:"user=dimuthu",
	}

	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)

}