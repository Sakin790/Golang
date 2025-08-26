package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8081")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	scanner := bufio.NewScanner(os.Stdin)
	serverScanner := bufio.NewScanner(conn)

	// Read server messages
	go func() {
		for serverScanner.Scan() {
			fmt.Println(serverScanner.Text())
		}
	}()

	// Input username first
	for {
		if scanner.Scan() {
			text := scanner.Text()
			fmt.Fprintln(conn, text)
			break
		}
	}

	fmt.Println("You joined the chat! Type /quit to exit.")

	// Send messages to server
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Fprintln(conn, text)
		if text == "/quit" {
			break
		}
	}
}
