package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main(){
	fmt.Println("welcome")
	performGetReq()
}

func performGetReq(){
	const myurl = "http://localhost:8000/get"
	res,err := http.Get(myurl)
	if err != nil{
		panic(err)
	}

	defer res.Body.Close()
	fmt.Println("status code",res.StatusCode)
	fmt.Println("content length is :",res.ContentLength)

	var responseString strings.Builder
	content,_ := ioutil.ReadAll(res.Body)
	byteCount,_ := responseString.Write(content)
	fmt.Println("Byte Count is :",byteCount)
	fmt.Println(responseString.String())

	//fmt.Println(content)
	//fmt.Println(stringContent)
}