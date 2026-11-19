package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

var clients = make(map[net.Conn]string) // connected clients

func main() {
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	fmt.Println("🚀 TCP chat server running on :8081")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Connection error:", err)
			continue
		}

		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)

	// Ask for username
	var name string
	for {
		conn.Write([]byte("Enter your username: "))
		if !scanner.Scan() {
			return
		}
		name = strings.TrimSpace(scanner.Text())
		if name != "" {
			break
		}
		conn.Write([]byte("Username cannot be empty.\n"))
	}

	// Save client and notify
	clients[conn] = name
	broadcast(fmt.Sprintf("%s joined the chat!", name), conn)

	// Read messages
	for scanner.Scan() {
		msg := scanner.Text()
		if strings.ToLower(msg) == "/quit" {
			broadcast(fmt.Sprintf("%s left the chat", name), conn)
			delete(clients, conn)
			break
		}
		broadcast(fmt.Sprintf("%s: %s", name, msg), conn)
	}
}

func broadcast(message string, sender net.Conn) {
	fmt.Println(message) // server console
	for conn := range clients {
		if conn != sender {
			conn.Write([]byte(message + "\n"))
		}
	}
}
