package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// First command: docker ps
	command := exec.Command("sudo", "docker", "ps")
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()
	if err != nil {
		fmt.Println("Error while running docker ps:", err)
	}

	// Second command: docker images
	nextCommand := exec.Command("sudo", "docker", "images")
	nextCommand.Stdout = os.Stdout
	nextCommand.Stderr = os.Stderr

	err = nextCommand.Run()
	if err != nil {
		fmt.Println("Error while running docker images:", err)
	}
}
