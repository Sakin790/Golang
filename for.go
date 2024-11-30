package main

import "fmt"

func main() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}
	for i := range 5 {
		fmt.Println("range", i)
	}
	fmt.Println()
}

