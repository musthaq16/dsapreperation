package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {

	sort.Ints(nums)

	res := [][]int{}

	for i := 0; i < len(nums)-2; i++ {

		left := i + 1
		right := len(nums) - 1
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for left < right {
			sum := nums[left] + nums[right] + nums[i]

			if sum == 0 {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				left++
				right--
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
			} else if sum > 0 {
				right--
			} else {
				left++
			}
		}
	}

	return res

}

func main() {
	vals := []int{-2, 0, 0, 1, 1, 2, 2}
	res := threeSum(vals)
	fmt.Println(res)
}
