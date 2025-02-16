package main

import (
	"fmt"
)

func vigEncrypt(text string, key string) string {
	if len(key) < len(text) {
		for i := 0; i < len(text)-len(key); i++ {
			key += string(key[i])
		}
	}
	if len(key) > len(text) {
		for i := 0; i < len(key)-len(text); i++ {
			key = key[:len(key)-1]
		}
	}
	var result string
	for i := 0; i < len(text); i++ {
		if text[i] >= 'A' && text[i] <= 'Z' {
			result += string((int(text[i])-65+int(key[i%len(key)])-65)%26 + 65)
		} else if text[i] >= 'a' && text[i] <= 'z' {
			result += string((int(text[i])-97+int(key[i%len(key)])-65)%26 + 97)
		} else {
			result += string(text[i])
		}
	}
	return result
}
func vigDecrypt(text string, key string) string {
	if len(key) < len(text) {
		for i := 0; i < len(text)-len(key); i++ {
			key += string(key[i])
		}
	}
	if len(key) > len(text) {
		for i := 0; i < len(key)-len(text); i++ {
			key = key[:len(key)-1]
		}
	}
	var result string
	for i := 0; i < len(text); i++ {
		if text[i] >= 'A' && text[i] <= 'Z' {
			result += string((int(text[i])-65-int(key[i%len(key)])+26)%26 + 65)
		} else if text[i] >= 'a' && text[i] <= 'z' {
			result += string((int(text[i])-97-int(key[i%len(key)])+26)%26 + 97)
		} else {
			result += string(text[i])
		}
	}
	return result
}
func vigenere() {
	var option int
	fmt.Println("1. Encrypt")
	fmt.Println("2. Decrypt")
	fmt.Print("Enter Option: ")
	fmt.Scanln(&option)
	if option == 1 {
		var text string
		var key string
		fmt.Print("Enter Text: ")
		fmt.Scanln(&text)
		fmt.Print("Enter Key: ")
		fmt.Scanln(&key)
		fmt.Println("Encrypted Text: ", vigEncrypt(text, key))
	} else if option == 2 {
		var text string
		var key string
		fmt.Print("Enter Text: ")
		fmt.Scanln(&text)
		fmt.Print("Enter Key: ")
		fmt.Scanln(&key)
		fmt.Println("Decrypted Text: ", vigDecrypt(text, key))
	} else {
		fmt.Println("Invalid Option")
	}
}
