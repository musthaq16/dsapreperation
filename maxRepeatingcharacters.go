package main

import "fmt"

func maxRepeatingCharacter(str string) string {

	hashMap := make(map[rune]int)

	for _, val := range str {
		hashMap[val]++
	}
	maximum := 0
	maxChar := ""
	for key, val := range hashMap {
		if val > maximum {
			maximum = val
			maxChar = string(key)
		}
	}
	return maxChar

}

func main() {
	inp := "aaaaabbbbbbbbbcddccccccc"

	res := maxRepeatingCharacter(inp)

	fmt.Println(res)
}
