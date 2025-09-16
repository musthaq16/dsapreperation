package main

import "fmt"

func longestCommonSubstring(val string) int {
	seen := make(map[byte]int)

	left := 0
	maximum := 0
	for right := 0; right < len(val); right++ {

		ch := val[right]

		if idx, exist := seen[ch]; exist && idx >= left {
			left = idx + 1
		}

		seen[ch] = right

		if right-left+1 > maximum {
			maximum = right - left + 1
		}

	}

	return maximum
}

func main() {
	val := "pwwkeuiw"

	max := longestCommonSubstring(val)

	fmt.Println(max)
}
