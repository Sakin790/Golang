package main
// var arrayName [size]Type
import (
	"fmt"
)

func main() {

	loop()
}

func array() {
	var myArr = [5]int{12, 3, 24, 52}
	fmt.Println(myArr)

	myArr2 := [...]int{12, 34, 13, 14}
	fmt.Println(myArr2)
	fmt.Println(len(myArr2))
	myArr2[2] = 14
	fmt.Println(myArr2[2])
}

func loop() {

	array := [...]int{10, 20, 30, 40, 50}
	for i := 0; i < len(array); i++ {
		fmt.Println(i, array[i])
	}

	name := [...]string{"messi", "cr7", "neymar", "Zidane"}
	for i := 0; i < len(name); i++ {
		fmt.Println(name[i])
	}

	club := [...]string{"City", "Liverpool", "United", "Madrid"}

	for index, value := range club {
		fmt.Println(index, value)
	}

}
