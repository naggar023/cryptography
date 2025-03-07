package main

import (
	"fmt"
)

func main() {
	var option int
	fmt.Println("1. Ceaser")
	fmt.Println("2. Vigenere")
	fmt.Println("3. Bruteforce attack on Monoalphabetic Cipher")
	fmt.Println("4. Cryptanalysis on Monoalphabetic Cipher")
	fmt.Println("5. Playfair Cipher")
	fmt.Print("Enter Option: ")
	fmt.Scanln(&option)
	switch option {
	case 1:
		ceaser()
	case 2:
		vigenere()
	case 3:
		task1()
	case 4:
		task2()
	case 5:
		task3()
	default:
		fmt.Println("Invalid Option")
	}
}
