package main

import "fmt"

func productArray(nums []int) []int {

	rightArr := make([]int, len(nums))
	leftArr := make([]int, len(nums))
	result := make([]int, len(nums))

	leftArr[0] = 1
	for i := 1; i < len(nums); i++ {
		leftArr[i] = leftArr[i-1] * nums[i-1]
	}

	rightArr[len(nums)-1] = 1
	for i := len(nums) - 2; i >= 0; i-- {
		rightArr[i] = rightArr[i+1] * nums[i+1]
	}

	for i := 0; i < len(nums); i++ {
		result[i] = leftArr[i] * rightArr[i]
	}

	return result

}

func main() {
	val := []int{-1, 1, 0, -3, 3}
	productArr := productArray(val)

	fmt.Println(productArr)

}
