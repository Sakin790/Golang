package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Enter your name: ")
	name, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(name)
	userName := []string{"sakin", "mahid", "abir", "Tamim"}
	if name == userName[0] {
		fmt.Println("Hello", name)
	}

}
