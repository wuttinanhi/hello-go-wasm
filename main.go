package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func SendGetRequest(url string) {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	text := string(body)
	fmt.Println(text)
}
func main() {
	fmt.Println("Hello, wasm")
	SendGetRequest("README.md")
}
