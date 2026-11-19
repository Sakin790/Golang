package main

import "fmt"

func rangeFunc() {
	var count int
	for range 10 {
		count++
		fmt.Println("Range function", count)
	}
}

func main() {
	var name string = "sakin"
	var a string
	fmt.Println(name)
	fmt.Println(a)

	var student string
	student = "Mahid"
	fmt.Println(student)

	age := 23
	fmt.Println(age)
	rangeFunc()
}
