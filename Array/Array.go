package main

import (
	"fmt"
)

func main() {

  var arr1 = [3]int{1,2,3}
  fmt.Println(arr1)

  
  arr2 := [5]int{4,5,6,7,8}
  fmt.Println(arr2)

  var infinite = [...]int{12, 34,55,53,432,13}
  fmt.Println(infinite)

  var cars = [4]string{"Volvo", "BMW", "Ford", "Mazda"}
  fmt.Print(cars[1])

}