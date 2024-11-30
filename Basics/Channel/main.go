package main

import (
	"fmt"
	"time"
)
func processNumber(numberChan chan int,) {
	fmt.Println("Processing number", <-numberChan)
}


func main() {
	numberChan := make(chan int) //creating channel
	go processNumber(numberChan)
	numberChan <- 5 //passing data to channel
	time.Sleep(time.Second*2) //hold main function for 2 second
}


