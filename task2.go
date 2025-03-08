package main

import (
	"fmt"
	"sort"
)

func task2() {
	var ct string
	fmt.Println("Enter Cipher Text Encrypted with Monoalphabetic cipher")
	fmt.Scanln(&ct)
	englishFreq := "etaoinshrdlcumwfgypbvkjxqz"
	ctFreq := make(map[rune]int)
	for _, letter := range ct {
		ctFreq[letter]++
	}

	ctLetters := make([]rune, 0, len(ctFreq))
	for letter := range ctFreq {
		ctLetters = append(ctLetters, letter)
	}

	sort.Slice(ctLetters, func(i, j int) bool {
		return ctFreq[ctLetters[i]] > ctFreq[ctLetters[j]] // Sort by frequency descending
	})

	mapping := make(map[rune]rune)
	for i, letter := range ctLetters {
		if i < len(englishFreq) {
			mapping[letter] = rune(englishFreq[i])
		}
	}

	decrypted := ""
	for _, letter := range ct {
		if mapped, exists := mapping[letter]; exists {
			decrypted += string(mapped)
		} else {
			decrypted += string(letter) // Keep original if not found
		}
	}

	fmt.Println("Mapping: ", mapping)
	fmt.Println("Decrypted Text: ", decrypted)

}
