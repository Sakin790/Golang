package main

import "fmt"

func main() {
	footBallClub := [3]string{
		"FCB",
		"RMZ",
		"NFO",
	}

	// Loop through the array with both index and value
	for index, value := range footBallClub {
		fmt.Println(index, value)
	}
}
