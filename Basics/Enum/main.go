package main

import "fmt"

type oderStatus string

const (
	Receving  oderStatus = "Receving"
	Confirm   oderStatus = "Confirm"
	Prepared  oderStatus = "Prepared"
	Delivered oderStatus = "Delivered"
)

func changeOderStatus(status oderStatus) {
	fmt.Println("Changing status to: ", status)

}
func main() {
	changeOderStatus(Confirm)
}
