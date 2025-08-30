package main

import "fmt"

func printGeneric[T int | string | bool](data T) (T, error) {
	fmt.Println("Received:", data)
	return data, nil
}
func main() {
	value, err := printGeneric(false)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Returned:", value)
	}
}
