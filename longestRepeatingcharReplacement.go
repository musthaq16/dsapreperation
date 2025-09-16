package main

import "fmt"

func characterReplacement(s string, k int) int {

	hashMap := make(map[byte]int)

	left := 0
	maxFreq := 0
	res := 0

	for right := 0; right < len(s); right++ {
		ch := s[right]
		hashMap[ch]++

		if hashMap[ch] > maxFreq {
			maxFreq = hashMap[ch]
		}

		//reseize the window if it pass
		if (right-left+1)-maxFreq > k {
			hashMap[ch]--
			left++
		}

		if right-left+1 > res {
			res = right - left + 1
		}
	}

	return res
}

func main() {
	val := "ABAB"
	k := 2
	res := characterReplacement(val, k)
	fmt.Println(res)

}
