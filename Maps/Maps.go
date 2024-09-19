package main

import "fmt"

func main() {
	footBall()
}

//like and object
func footBall() {

	playerRating := map[string]int{
		"messi":  10,
		"neymar": 8,
		"cr7":    9,
		"r9":     9,
	}
	fmt.Println(playerRating["messi"])
	playerRating["Foden"] = 3
	fmt.Println(playerRating["foden"])

	// Loop er jonno Range use korbo
	for key, value := range playerRating {
		fmt.Printf("%s : %d\n", key, value)
	}

	_, exists := playerRating["messis"]
	if exists {
		fmt.Println("Messi exist")
	} else {
		fmt.Println("Not Exists")
	}
}

func practice1() {
	m := map[string]int{
		"apple":  1,
		"banana": 2,
	}
	m["mango"] = 3 //adding value
	fmt.Print(m["mango"])
	fmt.Println(m["Sakin"]) // empty value

	value, exists := m["apple"]
	if exists {
		fmt.Println("Apple exist in Maps", value)
	} else {
		fmt.Println("Dose not Exist")
	}

	//delete(m, "mango")

	for key, value := range m {
		fmt.Printf("%s ** %d\n", key, value)
	}
}
