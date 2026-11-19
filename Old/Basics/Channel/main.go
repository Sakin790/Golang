package main

import (
	"fmt"
	"time"
)

func baker(cakes chan string) {

	for i := 0; i < 5; i++ {

		cake := fmt.Sprintf("Cake#%d", i)
		fmt.Println("Baker produced:", cake)
		cakes <- cake
		time.Sleep(time.Second)
	}
	close(cakes)

}

func customer(cakes chan string) {
	for cake := range cakes {
		fmt.Println("Customer bought:", cake)
		time.Sleep(2 * time.Second)
	}
}

func main() {

}
