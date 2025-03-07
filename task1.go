package main

import (
	"fmt"
)

func keyGen(str string, l, r int, ct string) {
	if l == r {
		decryptMono(str, ct)
		return
	}
	chars := []rune(str)
	for i := l; i <= r; i++ {
		chars[l], chars[i] = chars[i], chars[l]
		keyGen(string(chars), l+1, r, ct)
		chars[l], chars[i] = chars[i], chars[l]
	}
}
func decryptMono(key string, ct string) {
	pt := ""
	for i := 0; i < len(ct); i++ {
		// every character in ct could be any character in pt
		pt = pt + string(key[i])
	}
	fmt.Println(pt)
}
func task1() {
	defaultKey := "abcdefghijklmnopqrstuvwxyz"
	var ct string
	fmt.Println("Enter Encrypted text")
	fmt.Scanln(&ct)
	fmt.Println("Wanna make it run faster but less accurate? (0 for yes 1 for no)")
	var anss int
	fmt.Scanln(&anss)
	if anss == 0 {
		fmt.Println("enter how many charachter the key should be (less than 5 for fast output)")
		var n int
		fmt.Scanln(&n)
		key := defaultKey[:n]
		keyGen(key, 0, n-1, ct)
	} else {
		fmt.Println("go make a coffee and wait while the algorithm try to decrypt ...")
		keyGen(defaultKey, 0, len(defaultKey)-1, ct)
	}
}
