package main

import "fmt"

func main() {
	fmt.Println("Start")
	panic("Something went wrong!")
	fmt.Println("End") // This will not be printed

	value := divide(12, 2)
	fmt.Println(value)
}

func divide(a, b int) int {
	if b == 0 {
		panic("division by zero")
	}
	return a / b
}

func createPanic() {
	fmt.Println(1)
	fmt.Println(2)
	panic("Im not Excuting anymore....")
	fmt.Println(3)
}


func recoverY()  {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	fmt.Println("Starting")
	panic("Something went wrong")
	fmt.Println("This will not be printed")
}