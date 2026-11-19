package main

import "fmt"

type Rectangle struct {
	Width, Height int
}

func (r Rectangle) Area() (int, error) {

	return r.Width * r.Height, nil
}

func main() {
	rect := Rectangle{10, 5}
	fmt.Println(rect.Area()) // 50
}
