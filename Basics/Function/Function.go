/*

func functionName(parameter1 type, parameter2 type) returnType {
    // Function body
    return value
}

*/

package main

import "fmt"

func main() {

	greet()
	result,err := toSum(2, 3)
	if err != nil {
	
	}
	fmt.Println(result)
	r, q := multiple(5, 5)
	fmt.Println(q, r)
	fmt.Println("Area of rectangle:", rectangleArea(5, 3))

}
func greet() {
	fmt.Println("Welcome...")
}

func toSum(a, b int) (int, error) {
	return a + b, nil
}

//Function with Maltiple value

func multiple(a int, b int) (int, int) {
	r := a / b
	q := a % b
	return r, q // return korle vaiable a hold korte hobe, jehutu 2 ta value return korche
	// tahole amake 2 ta variable  nite hobee
	//variable a hold na korte hole , function return type a nam diye dite hoi
}

func rectangleArea(length, width int) (area int) {
	area = length * width
	return // The named return 'area' is automatically returned
}
