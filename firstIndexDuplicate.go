package main

import "fmt"

func firstIndexDuplicate(str string) int {

	hashMap := make(map[rune]int)

	for _, val := range str {
		hashMap[val]++
	}
	fmt.Println(hashMap)

	for index, val := range str {
		if hashMap[val] > 1 {
			return index
		}
	}

	return -1

}

func main() {

	str := "apple"
	// str := "àbà"

	res := firstIndexDuplicate(str)
	fmt.Println(res)
}
