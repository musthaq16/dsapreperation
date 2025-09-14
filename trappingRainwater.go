package main

import "fmt"

func trappingRainWater(nums []int) int {

	leftArr := make([]int, len(nums))
	rightArr := make([]int, len(nums))
	res := 0
	leftMax := 0
	rightMax := 0
	fmt.Println(len(nums))

	for i := 0; i < len(nums); i++ {
		if nums[i] > leftMax {
			leftMax = nums[i]
		}
		leftArr[i] = leftMax
	}

	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] > rightMax {
			rightMax = nums[i]
		}
		rightArr[i] = rightMax
	}

	for i := 0; i < len(nums); i++ {
		res = (min(leftArr[i], rightArr[i]) - nums[i]) + res
	}

	return res
}

func main() {
	vals := []int{4, 2, 0, 3, 2, 5}

	res := trappingRainWater(vals)
	fmt.Println(res)
}
