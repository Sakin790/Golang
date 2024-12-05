package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	IsAdult bool   `json:"is_adult"`
}

func main() {
	fmt.Println("Starting json converter")
	persons := Person{Name: "John Doe", Age: 25, IsAdult: true}
	data, err := json.Marshal(persons)
	if err != nil {
		fmt.Println("Error marshaling json:", err)
		return
	}
	fmt.Println(string(data))

	var personData Person

	err = json.Unmarshal(data, &personData)
	if err != nil {
		fmt.Println("json.Unmarshal error:", err)
		return
	}

	fmt.Println(personData)

}
