package main

import (
	"fmt"
	"time"
)

func function1() {
	fmt.Println("I am function 1")
}
func function2() {
	fmt.Println("I am function 2")
	time.Sleep(5000 * time.Millisecond)
}
func function3() {
	fmt.Println("I am function 3")
}

func loop() {
	for i := 0; i < 100; i++ {
		go fmt.Println("The Number is:", i)
	}
}

func main() {
	function1()
	function3()
	go loop()
	time.Sleep(10000 * time.Millisecond)
}
