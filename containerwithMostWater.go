package main

import "fmt"

func containerWithMostWater(nums []int) int {

	left := 0
	right := len(nums) - 1
	maximum := 0

	for left <= right {
		width := right - left
		height := min(nums[left], nums[right])
		maxval := width * height
		maximum = max(maximum, maxval)
		if nums[left] < nums[right] {
			left++
		} else {
			right--
		}

	}
	return maximum
}

func main() {
	vals := []int{1, 1}

	max := containerWithMostWater(vals)
	fmt.Println(max)
}
