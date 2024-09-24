package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Time is Everthing")

	now := time.Now()
	fmt.Println("Default format:", now)
	fmt.Println("Custom format:", now.Format("2006-01-02 15:04:05"))
	fmt.Println("Just the time:", now.Format("15:04:05"))
	fmt.Println("Just the date:", now.Format("02-Jan-2006"))

	timer := time.NewTimer(2 * time.Second)
	fmt.Println("Waiting for the timer...")

	<-timer.C // Blocks until the timer expires
	fmt.Println("Timer expired!")

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for i := 0; i < 3; i++ {
		<-ticker.C
		fmt.Println("Tick at", time.Now())
	}
	interval()
}

func interval() {
	// Create a new ticker that ticks every 2 seconds
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop() // Ensure the ticker is stopped to avoid resource leaks
	// Loop 10 times, printing "Hello, World!" every 2 seconds
	for i := 0; i < 10; i++ {
		<-ticker.C // Block until the ticker sends a value
		fmt.Println("Hello, World!")
	}
}
