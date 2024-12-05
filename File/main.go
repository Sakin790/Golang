package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	/*
		fmt.Println("Creating File...")
		file, err := os.Create("Example.txt")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()

		content := "Hello, world!"
		byte, err := io.WriteString(file, content+"\n")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(byte)
		fmt.Printf("Type of file is %T\n", file)
	*/
	file, err := os.Open("Example.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	//make a buffer for reading the file

	buffer := make([]byte, 1024)
	for {
		n, err := file.Read(buffer)
		fmt.Println(n)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			return
		}
		// Reading data from buffer
		fmt.Println(string(buffer[:n]))
	}

}
