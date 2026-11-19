package main

import "fmt"

type students struct {
	id    string
	name  string
	age   int
	class string
}

func (pointer *students) changeGender(name string) {
	pointer.name = name
}
func main() {

	sakin := students{
		id:    "001",
		name:  "Sakin",
		age:   23,
		class: "10th",
	}
	fmt.Println(sakin.age)
	sakin.changeGender("Mahid")
	fmt.Println(sakin.name)

}
