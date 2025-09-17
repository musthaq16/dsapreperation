package main

import (
	"fmt"
	"sort"
)

func findMissingNumber(nums []int) int {

	sort.Ints(nums)

	fmt.Println(nums)

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1]+1 {
			return nums[i-1] + 1
		}
	}

	return -1

}

func main() {
	nums := []int{55, 57, 56, 60, 58}
	// 55,56,57,58,60

	res := findMissingNumber(nums)

	fmt.Println(res)
}
