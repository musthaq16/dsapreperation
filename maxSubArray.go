package main

import "fmt"

func subArray(nums []int) int {
	currSum := nums[0]
	maxSum := nums[0]

	for i := 1; i < len(nums); i++ {
		currSum = max(nums[i], nums[i]+currSum)
		maxSum = max(maxSum, currSum)
	}
	return maxSum
}

func main() {
	vals := []int{5, 4, -1, 7, 8}
	val := subArray(vals)
	fmt.Println(val)
}
