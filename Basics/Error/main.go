package main

import (
	"errors"
	"fmt"
)

func devided(a, b float64) (float64, error) {

	if b == 0 {
		return 0, errors.New("cannot devided bt zero")
	}
	return a / b, nil
}

func main() {

	data, error := devided(10, 2)
	if error != nil {
		fmt.Println(error)
	}
	fmt.Println(data)

}
