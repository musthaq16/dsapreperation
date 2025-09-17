package main

import "fmt"

func findMissingNaturalNumber(nums []int) int {

	n := len(nums)

	excpected := n * (n + 1) / 2

	actual := 0

	for i := 0; i < n; i++ {
		actual = actual + nums[i]
	}

	return excpected - actual

}

func main() {

	inp := []int{9, 6, 4, 2, 3, 5, 7, 0, 1}

	res := findMissingNaturalNumber(inp)

	fmt.Println(res)
}
