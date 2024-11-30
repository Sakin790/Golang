package main

import "fmt"

/*
format speciflers


%d\n => int
%s\n => string
%.3f =>float
%T => type
*/

func main() {
	age := 23
	name := "Sakin"
	balance := 23.78654
	fmt.Println(name, age, balance)

	//printf => print with custom format
	fmt.Printf("My name is %s\n", name)
	fmt.Printf("My gae is %d\n", age)
	fmt.Printf("My name is %.3f\n", balance)

	//Type check
	fmt.Printf("The type of age is %T\n", age)
	fmt.Printf("The type of name is %T\n", name)
	fmt.Printf("The type of balance is %T\n", balance)

}
