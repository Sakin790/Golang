package main

import (
	"fmt"
	"strings"
)

func main() {
	data := "apple, banana, mango"
	fmt.Println(data)
	parts := strings.Split(data, ",")
	fmt.Printf("The type of parts %T\n", parts)
	fmt.Println(parts)
}
