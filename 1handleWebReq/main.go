package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)
const url = "https://lco.dev"

func main() {
	fmt.Println("web req")
	//check type response
	res,err := http.Get(url)

	if err != nil{
		panic(err)
	}

	fmt.Printf("Response is of type: %T\n",res)
	//always need to cos once done
	defer res.Body.Close() 

	data,err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}
	//we hve to convert data into string
	content := string(data)
	fmt.Printf(content)
}