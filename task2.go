package main

import (
	"fmt"
	"sort"
)

func task2() {
	var ct string
	fmt.Println("Enter Cipher Text Encrypted with Monoalphabetic cipher")
	fmt.Scanln(&ct)
	englishFreq := []string{
		"etaoinshrdlcumwfgypbvkjxqz",
		"taeoinshrdlcumwfgypbvkjxqz",
		"oetaoinshrdlcumwfgypbvkjxqz",
		"ineotasrdlcumwfgypbvkjxqz",
		"soetainhrdlcumwfgypbvkjxqz",
	}
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

	for i, englishFreq := range englishFreq {
		mapping := make(map[rune]rune)
		for j, letter := range ctLetters {
			if j < len(englishFreq) {
				mapping[letter] = rune(englishFreq[j])
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

		fmt.Printf("Probable Decryption %d: %s\n", i+1, decrypted)
	}
}
