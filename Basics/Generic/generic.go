package main

import (
	"fmt"
)

// function for int
func printInt(number []int) {
	for _, value := range number {
		fmt.Println(value)
	}

}

// function for strings
func printString(names []string) {

	for _, value := range names {
		fmt.Println(value)
	}

}



// generic function
func generic[T int | string | bool](item []T) {

	for _, value := range item {
		fmt.Println(value)
	}
}

//generic for structs

type stack[T int | string | bool] struct {
	element []T
}

func main() {

	myStack := stack[string]{
		element: []string{"Yesssssssssssss"},
	}
	fmt.Println(myStack)

	printInt([]int{1, 3, 3, 4})
	printString([]string{"Sakin", "Tamim", "Mahid"})

	//socre := []int{12, 2, 34, 21}
	name := []string{"Messsi"}
	generic(name)
}
