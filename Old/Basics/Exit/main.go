package main

import (
	"fmt"
	"log"
)

func main() {
	defer fmt.Println("Im deffer Function") // os.Exit(0) never allow this
	fmt.Println("Program started")
	//os.exit(0)=== to log Fatal 
	log.Fatal("Fatal error, program will exit") // Exits the program with a status code of 0
	fmt.Println("This will not be printed")
}
