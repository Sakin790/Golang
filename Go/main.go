package main

import (
	"fmt"
	"time"
)

func task(id int) {
	fmt.Println("Doing task", id)
}

func main() {
	for i := 0; i <= 10; i++ {
		go func(id int) {
			task(id)
		}(i)
		fmt.Println()
	}

	time.Sleep(time.Second * 2)
}
