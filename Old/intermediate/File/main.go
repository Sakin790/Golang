package main

import (
	"fmt"
	"os"
)

func main() {
	readFile()
}

func write() {
	data := "Lets have some fun with golang"
	err := os.WriteFile("note.txt", []byte(data), 0644)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Write file Done")
	}
}

func create() {
	file, err := os.Create("note.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fmt.Println("File creared")
}

func readFile() {
	data, err := os.ReadFile("note.txt")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("File data is:", string(data))
	}
}

//Folder


