package main

import (
	"fmt"
	"strings"
	"unicode"
)

func isSentencePalindrome(strs string) bool {

	str := strings.ToLower(strs)

	filtered := []rune{}

	for _, val := range str {
		if unicode.IsLetter(val) {
			filtered = append(filtered, val)
		}
	}

	left := 0
	right := len(filtered) - 1
	for left < right {
		if filtered[left] != filtered[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func main() {
	fmt.Println(isSentencePalindrome("A man, a plan, a canal: Panama")) // true
	fmt.Println(isSentencePalindrome("race a car"))                     // false
	fmt.Println(isSentencePalindrome("No lemon, no melon"))
}
