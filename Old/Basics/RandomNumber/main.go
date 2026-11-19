package main

import (
	"fmt"
	"math/rand"
)

func main() {
	min := 10
	max := 50

	// Generate a random number in the range [min, max)
	randomNumber := rand.Intn(max-min) + min
	fmt.Println(randomNumber)
}
