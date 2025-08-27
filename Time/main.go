
package main

import (
	"fmt"
	"time"
)

func printHello() {
	fmt.Println("Printing Hello")
}

func printHelloTwo() {
	fmt.Println("Welcome")
}
func main() {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop() // stop ticker when program exits

	for {
		<-ticker.C // wait for next tick
		printHello()
		printHelloTwo()
	}
}
