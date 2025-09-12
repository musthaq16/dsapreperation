package main

import "fmt"

func findDuplicates(nums []int) int {

	slow := nums[0]
	fast := nums[0]

	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			fmt.Println("vall iss", slow) // first find intersection. then reset slow and both move 1 step.
			break
		}
	}
	slow = nums[0]
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}

	return slow

}

func main() {
	dup := []int{3, 1, 3, 4, 2}
	res := findDuplicates(dup)

	fmt.Println(res)

}
