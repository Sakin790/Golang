package main

import (
	"fmt"
	"os/exec"
	"strings"
	"sync"
)

func main() {
	out, err := exec.Command("sudo", "docker", "images").Output()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	lines := strings.Split(string(out), "\n")[1:] // skip header
	var wg sync.WaitGroup

	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) < 5 {
			continue
		}

		repo, tag, sizeStr := fields[0], fields[1], fields[4]
		wg.Add(1)

		go func(repo, tag, sizeStr string) {
			defer wg.Done()

			if repo == "mysql" && tag == "8" {
				fmt.Println("latest mysql")
			} else if parseSize(sizeStr) > 1024 { // size in MB
				fmt.Println(repo, tag, "Too big size")
			} else {
				fmt.Println(repo, tag, "container is perfect")
			}
		}(repo, tag, sizeStr)
	}

	wg.Wait()
}

// parseSize converts size string like "786MB" or "1.5GB" to MB
func parseSize(s string) float64 {
	var size float64
	var unit string
	fmt.Sscanf(s, "%f%s", &size, &unit)
	if unit == "GB" {
		size *= 1024
	}
	return size
}
