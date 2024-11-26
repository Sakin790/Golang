package main

import "fmt"

// Bubble Sort function to sort an array
func bubbleSort(arr []int) {
	n := len(arr)

	// Perform sorting
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			// Swap if the element found is greater than the next element
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j] // Swap the elements
			}
		}
	}
}

func main() {
	// Declare and initialize the array
	arr := []int{50, 20, 40, 10, 234}

	// Call the bubbleSort function
	bubbleSort(arr)

	// Print the sorted array
	fmt.Println("Sorted array:", arr)
}
