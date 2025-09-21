package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) Greet() {
	fmt.Println(p.Name, p.Age)

}

type Employe struct {
	Person
	company string
}

func (e Employe) work() {

	fmt.Println(e.company)

}
func main() {

	idx := Person{
		Name: "sakin",
		Age:  12,
	}

	idx2 := Employe{
		company: "Google",
	}

	idx2.work()
	fmt.Println(idx.Name)
	idx.Greet()

	data := map[string]string{
		"sakin": "islam",
	}
	fmt.Println(data["sakin"])
}
