package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"syscall/js"
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

func Sum(a, b int) int {
	return a + b
}

func main() {
	hold := make(chan bool)

	fmt.Println("Hello, wasm")
	SendGetRequest("README.md")

	SumWrapper := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return Sum(args[0].Int(), args[1].Int())
	})

	js.Global().Set("sum", SumWrapper)

	<-hold
}
