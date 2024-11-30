package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// var name string
	// fmt.Scan(&name)
	// fmt.Println("Hello", name)
	fmt.Printf("Enter your name: ")
	name, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(name)
}


