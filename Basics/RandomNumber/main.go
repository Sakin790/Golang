package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Define the range (e.g., 10 to 50)
	min := 10
	max := 50

	// Generate a random number in the range [min, max)
	randomNumber := rand.Intn(max-min) + min
	fmt.Println(randomNumber)
}
