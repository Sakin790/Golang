package main

import "fmt"

type Vehicle interface {
	Start() (string, error)
	Stop()
}
type Car struct{}

func (c Car) Start() (string, error) {
	return "Car Has start", nil

}

func (c Car) Stop() {
	fmt.Println("Car Has Stop")

}

func main() {

}
