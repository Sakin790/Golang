package main

import "fmt"

func main() {

}

func buffer() {

	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	ch <- 3 //blocked, buffer is full
	// Buffered channels: sender only blocks when buffer is full, receiver blocks when buffer is empty.
}

func unbuffer() {
	ch := make(chan int)

	go func() {
		ch <- 291111
	}()

	recived := <-ch
	fmt.Println(recived)
}
