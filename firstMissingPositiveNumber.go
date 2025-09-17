package main

import "fmt"

func firstMissingPositive(arr []int) int {

	for i := 0; i < len(arr); {
		correctIndex := arr[i] - 1

		if arr[i] > 0 && arr[i] <= len(arr) && arr[i] != arr[correctIndex] {
			arr[i], arr[correctIndex] = arr[correctIndex], arr[i]
		} else {
			i++
		}
	}

	for i := 0; i < len(arr); i++ {
		if arr[i] != i+1 {
			return i + 1
		}
	} 

	return len(arr) + 1
}

func main() {
	fmt.Println(firstMissingPositive([]int{3, 4, -1, 1}))     // 2
	fmt.Println(firstMissingPositive([]int{1, 2, 0}))         // 3
	fmt.Println(firstMissingPositive([]int{7, 8, 9, 11, 12})) // 1
}
