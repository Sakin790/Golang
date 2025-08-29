package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	encodingJson()
}

func encodingJson() {

	type Person struct {
		Name    string `json:"name"`
		Age     int    `json:"age"`
		IsAdult bool   `json:"is_adult"`
	}

	persosn := Person{
		Name:    "sakin",
		Age:     22,
		IsAdult: true,
	}

	data, err := json.Marshal(persosn)
	if err != nil {
		fmt.Println("error while marshalling:", err)

	}

	fmt.Println(string(data))

	var personData Person

	err = json.Unmarshal(data, &personData)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(personData)
	}

}

func constant() {
	const name string = "sakin"

	const (
		port = 8080
		host = "localhost"
	)
	fmt.Println(name, port, host)
}

func arrayFunc() {
	score := [...]int{10, 20, 30, 40, 50}
	//score.append(score, 60) not dinamice resizable
	capacity := cap(score)
	length := len(score)
	fmt.Println(length, capacity)

	//slice
	mark := []int{100, 200, 300}
	mark = append(mark, 400) // can possible
	fmt.Println(mark)

	for i := 0; i < len(mark); i++ {
		fmt.Println(mark[i])

	}
	for index, value := range mark {
		fmt.Println(index, value)
	}
}

func rangeFunction() {
	for range 1000000 {
		fmt.Println("Hello Sakin")
	}
}
