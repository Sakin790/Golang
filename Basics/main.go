// sliceName := []datatype { value1, value2, value3,...valueN }

package main

import "fmt"

func main() {
	name := []string{"Leo", "Messi", "CR4", "CR5", "CR6", "CR7", "CR8", "CR9"}
	for i := 0; i < len(name); i++ {
		fmt.Print(" ", name[i])
	}

	for index, value := range name {
		fmt.Println("Index:", index, "Value:", value)
	}
}
