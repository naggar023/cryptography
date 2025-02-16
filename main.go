package main

import (
	"fmt"
)

func main() {
	var option int
	fmt.Println("1. Ceaser")
	fmt.Println("2. Vigenere")
	fmt.Print("Enter Option: ")
	fmt.Scanln(&option)
	switch option {
	case 1:
		ceaser()
	case 2:
		vigenere()
	default:
		fmt.Println("Invalid Option")
	}
}
