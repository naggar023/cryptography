package main

import (
	"fmt"
	"strings"
)

func task3() {
	var key string
	fmt.Println("Enter The key")
	fmt.Scanln(&key)

	var pfMat [5][5]rune
	pfarr := []rune("abcdefghiklmnopqrstuvwxyz") // Without j
	uniqueKey := ""
	seen := make(map[rune]bool)

	for _, ch := range key {
		ch = rune(strings.ToLower(string(ch))[0]) // Convert to lowercase
		if ch == 'j' {
			ch = 'i' // convert j to i
		}
		if !seen[ch] && strings.ContainsRune(string(pfarr), ch) {
			uniqueKey += string(ch)
			seen[ch] = true
		}
	}

	newPfarr := uniqueKey
	for _, ch := range pfarr {
		if !seen[ch] {
			newPfarr += string(ch)
			seen[ch] = true
		}
	}

	k := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			pfMat[i][j] = rune(newPfarr[k])
			k++
		}
	}

	fmt.Println("1. Encrypt")
	fmt.Println("2. Decrypt")
	var op int
	fmt.Scanln(&op)

	if op == 1 {
		var pt string
		fmt.Println("Enter text to Encrypt")
		fmt.Scanln(&pt)
		pt = strings.ReplaceAll(strings.ToLower(pt), "j", "i") // Normalize input
		if len(pt)%2 != 0 {
			pt += "x" // Padding if needed
		}

		ans := ""
		for i := 0; i < len(pt); i += 2 {
			r1, c1 := findPosition(pfMat, rune(pt[i]))
			r2, c2 := findPosition(pfMat, rune(pt[i+1]))

			if r1 == r2 {
				ans += string(pfMat[r1][(c1+1)%5]) // move right
				ans += string(pfMat[r2][(c2+1)%5])
			} else if c1 == c2 {
				ans += string(pfMat[(r1+1)%5][c1]) // move down
				ans += string(pfMat[(r2+1)%5][c2])
			} else {
				ans += string(pfMat[r1][c2])
				ans += string(pfMat[r2][c1])
			}
		}
		fmt.Println("Encrypted Text:", ans)
	} else if op == 2 {
		var ct string
		fmt.Println("Enter Encrypted text to Decrypt")
		fmt.Scanln(&ct)
		ct = strings.ToLower(ct)
		ans := ""
		for i := 0; i < len(ct); i += 2 {
			r1, c1 := findPosition(pfMat, rune(ct[i]))
			r2, c2 := findPosition(pfMat, rune(ct[i+1]))

			if r1 == r2 {
				ans += string(pfMat[r1][(c1+4)%5]) // Move left in the row
				ans += string(pfMat[r2][(c2+4)%5])
			} else if c1 == c2 {
				ans += string(pfMat[(r1+4)%5][c1]) // Move up in the column
				ans += string(pfMat[(r2+4)%5][c2])
			} else {
				ans += string(pfMat[r1][c2])
				ans += string(pfMat[r2][c1])
			}
		}
		fmt.Println("Decrypted Text:", ans)
	} else {
		fmt.Println("Invalid option")
	}
}

func findPosition(mat [5][5]rune, ch rune) (int, int) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if mat[i][j] == ch {
				return i, j
			}
		}
	}
	return -1, -1 // should never happen
}
