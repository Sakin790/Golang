package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(2 * time.Second)
	fmt.Println("Waiting for the timer...")

	<-timer.C  // Blocks until the timer expires
	fmt.Println("Timer expired!")
}
