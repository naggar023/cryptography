package main

import (
	"fmt"
)

func encrypt(text string, shift int) string {
	var result string
	for i := 0; i < len(text); i++ {
		if text[i] >= 'A' && text[i] <= 'Z' {
			result += string((int(text[i])-65+shift)%26 + 65)
		} else if text[i] >= 'a' && text[i] <= 'z' {
			result += string((int(text[i])-97+shift)%26 + 97)
		} else {
			result += string(text[i])
		}
	}
	return result
}
func decrypt(text string, shift int) string {
	var result string
	for i := 0; i < len(text); i++ {
		if text[i] >= 'A' && text[i] <= 'Z' {
			result += string((int(text[i])-65-shift+26)%26 + 65)
		} else if text[i] >= 'a' && text[i] <= 'z' {
			result += string((int(text[i])-97-shift+26)%26 + 97)
		} else {
			result += string(text[i])
		}
	}
	return result
}
func main() {
	var option int
	fmt.Println("1. Encrypt")
	fmt.Println("2. Decrypt")
	fmt.Print("Enter Option: ")
	fmt.Scanln(&option)
	if option == 1 {
		var text string
		var shift int
		fmt.Print("Enter Text: ")
		fmt.Scanln(&text)
		fmt.Print("Enter Shift: ")
		fmt.Scanln(&shift)
		fmt.Println("Encrypted Text: ", encrypt(text, shift))
	} else if option == 2 {
		var text string
		var shift int
		fmt.Print("Enter Text: ")
		fmt.Scanln(&text)
		fmt.Print("Enter Shift: ")
		fmt.Scanln(&shift)
		fmt.Println("Decrypted Text: ", decrypt(text, shift))
	} else {
		fmt.Println("Invalid Option")
	}
}
