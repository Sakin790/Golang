package main

import "fmt"

func main() {
	var num int = 23
	fmt.Printf("%T\n", num)

	var data float32 = float32(num)
	fmt.Printf("%T\n", data)

	var nums int = 50
	var numsToSting string = string(nums)
	fmt.Printf("%T\n", numsToSting) 


} 
