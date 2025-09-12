package main

import "fmt"

func twoSum(values []int, target int) []int {

	twoSumMap := make(map[int]int)

	for i, val := range values {
		comp := target - val
		if sum, exist := twoSumMap[comp]; exist {
			return []int{sum, i}
		}
		twoSumMap[val] = i
	}

	return []int{}
}

func main() {
	values := []int{3, 2, 4}
	target := 6
	result := twoSum(values, target)

	fmt.Println(result)
}
