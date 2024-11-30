package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("Web request")
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos")

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer res.Body.Close()
	fmt.Printf("type of response body : %T\n", res)

	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response body: ", err)
		return
	}
	fmt.Println(string(data))
}
