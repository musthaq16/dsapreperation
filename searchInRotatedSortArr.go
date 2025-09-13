package main

import "fmt"

func searchInRotSortArray(nums []int, target int) int {

	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (left + right) / 2

		if nums[mid] == target {
			return nums[mid]
		}

		if nums[mid] >= nums[left] { // left sub arr sorted
			if nums[left] <= target && nums[mid] > target {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else { // right sub arr sorted
			if nums[right] >= target && nums[mid] < target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return -1
}

func main() {
	vals := []int{4, 5, 6, 7, 0, 1, 2}
	target := 8
	foundedVal := searchInRotSortArray(vals, target)
	fmt.Println(foundedVal)
}
