package main

import "fmt"

func minElem(nums []int) int {
	left := 0
	right := len(nums) - 1
	for left < right {
		if nums[left] < nums[right] {
			return nums[left]
		}
		mid := (left + right) / 2
		fmt.Println("middd", mid)

		if nums[mid] >= nums[left] {
			left = mid + 1
		} else {
			right = mid
		}

	}

	return nums[left]
}

func main() {
	vals := []int{3, 4, 5, 1, 2}

	min := minElem(vals)
	fmt.Println(min)
}
