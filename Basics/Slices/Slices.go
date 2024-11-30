package main

import "fmt"

// sliceName := []datatype { value1, value2, value3,...valueN }
// mySlice := make ([]type , length [,capacity])
// slice without specifying the capacity
// mySlice2 := make([]int, 4)

//Dynamic Array
//uninitailize slices is nill
func main() {
	var nums []int
	fmt.Println(nums)
	fmt.Println(cap(nums))
	fmt.Println(len(nums))

	// slice initialize but no value
	var numbers = make([]int, 0, 5) // 0 length, 5 is capacity

	numbers = append(numbers, 10, 20, 30, 40, 50, 55)

	fmt.Println(numbers)
	fmt.Println(cap(numbers))

	//Other Methood
	score := []int{10, 20, 45, 44, 33} //main
	var myArr = [...]int{12, 3, 24, 52}
	fmt.Println(myArr)
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

	homeWork()
}

func homeWork() {
	work := []string{"Cooking", "Cleaning", "Watering"}
	fmt.Println(work)
	fmt.Println(work[0])
	fmt.Println(work[1])
	fmt.Println(work[2])
	fmt.Printf("The Length of  arrar %T\n", work)
}

func looping() {
	name := []string{"Leo", "Messi", "CR4", "CR5", "CR6", "CR7", "CR8", "CR9"}
	for i := 0; i < len(name); i++ {
		fmt.Print(" ", name[i])
	}

	for index, value := range name {
		fmt.Println("Index:", index, "Value:", value)
	}
}
