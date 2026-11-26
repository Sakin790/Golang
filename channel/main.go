package main

import "fmt"

type User struct {
	ID   int
	Name string
}

func main() {
	fmt.Println("Go channel")

	sakin := User{
		ID:   1,
		Name: "Sakin",
	}

	ch := make(chan User)

	go func() {
		ch <- sakin
	}()

	msg := <-ch
	fmt.Println(msg)
}
