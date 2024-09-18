package main

import "fmt"

func main() {

	var fuirt string
	fmt.Print("Enter your Fuirt name: ")
	fmt.Scan(&fuirt)
	var fuirtOption = []string{"apple", "mango", "Banana", "Guava"}
	switch fuirt {
	case fuirtOption[0]:
		fmt.Println("This is a Apple")
		break
	case fuirtOption[1]:
		fmt.Println("This is a mango")
		break
	case fuirtOption[2]:
		fmt.Println("Banana")
	case fuirtOption[3]:
		fmt.Println("Guava")
	default:
		fmt.Println("None of Them")
	}

}
