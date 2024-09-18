package main

import (
	"fmt"
)

func main() {
  var a, b, c, d int = 1, 3, 5, 7

  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
  fmt.Println(d)
  main2()
}




func main2() {
   var (
     a int
     b int = 1
     c string = "hello"
   )
  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
}