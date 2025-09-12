package main

import "fmt"

func maxProductSubarray(nums []int) int {

	maxSoFar := nums[0]
	minSoFar := nums[0]
	res := nums[0]

	for i := 1; i < len(nums); i++ {
		curr := nums[i]
		if curr < 0 {
			maxSoFar, minSoFar = minSoFar, maxSoFar
		}

		maxSoFar = max(curr, maxSoFar*curr)
		minSoFar = min(curr, minSoFar*curr)

		res = max(res, maxSoFar)
	}
	return res

}

func main() {

	nums := []int{-2, 1, -1}
	maxProduct := maxProductSubarray(nums)
	fmt.Println(maxProduct)
}
