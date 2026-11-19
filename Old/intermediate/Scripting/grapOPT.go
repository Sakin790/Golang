package main

import (
	"fmt"
	"os/exec"
	"time"
)

func main() {
	// Run "sudo docker images"
	cmd := exec.Command("sudo", "docker", "images")
	// Capture stdout + stderr together
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error while running docker images:", err)
		return
	}

	// Print as string
	fmt.Println("Output from docker images:")
	fmt.Println(string(out))
}
