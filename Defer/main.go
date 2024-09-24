package main

import "fmt"

func main() {
	fmt.Println("Starting")
	defer fmt.Println("Middle")
	fmt.Println("End")
}
