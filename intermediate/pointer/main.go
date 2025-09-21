package main

import "fmt"

// Function that uses a pointer
func updateValue(ptr *int) {
    *ptr = 50
}

func main() {
    x := 10
    fmt.Println("Before:", x) // Prints 10

    updateValue(&x) // Pass the address of x
    fmt.Println("After:", x) // Prints 50
}
