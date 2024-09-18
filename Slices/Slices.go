package main

import "fmt"

//Dynamic Array
//uninitailize slices is nill
func main() {
	var nums []int
	fmt.Println(nums)
	fmt.Println(cap(nums))
	fmt.Println(len(nums))

	//I we want not nill value
	var numbers = make([]int, 0, 5) //5 is capacity

	numbers = append(numbers, 10, 20, 30, 40, 50, 55)

	fmt.Println(numbers)
	fmt.Println(cap(numbers))

	//Other Methood
	score := []int{10, 20, 45, 44, 33}
	fmt.Println(score)
	fmt.Println(cap(score))



	
	var Football = make([]string, 1, 5)
	Football = append(Football, "UFA", "PL", "EPL")
	var Cricket = make([]string, len(Football), 5)
	copy(Cricket, Football)
	fmt.Println(Cricket)

	var emp = make([]string, 1, 5)
	emp = append(emp, "sakin", "mahid", "abir")
	var cmp = make([]string, len(emp))
	copy(cmp, emp)
	fmt.Println(cmp)

}
