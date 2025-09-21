package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	data := "apple, banana, mango"
	fmt.Println(data)
	parts := strings.Split(data, ",")
	fmt.Printf("The type of parts %T\n", parts)
	fmt.Println(parts)

	str := "123"
	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Error to conversion:", err)
		return
	} else {
		fmt.Println("Integer Data", num)
	}

}
