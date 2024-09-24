package main

import (
	"log"
	"os"
)

func main() {
	// Open a file for logging
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Set output of logs to file
	log.SetOutput(file)

	// Log messages
	log.Println("Logging to a file")
}
