package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func any() {
	var name string
	fmt.Println("Enter Your Name: ")
	fmt.Scanln(&name)
	fmt.Println("Your Name is", name)
}

func alternative() {
	idx := []string{} // empty slice

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Something: ")
	data, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	words := strings.Fields(data)
	idx = append(idx, words...)

	for i := 0; i < len(idx); i++ {
		fmt.Println(idx[i])
	}

}

func main() {
	alternative()
}
