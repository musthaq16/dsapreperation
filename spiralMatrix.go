package main

import "fmt"

func printSpiralMatrix(matrix [][]int) []int {
	res := make([]int, 0)

	top := 0
	bottom := len(matrix) - 1
	left := 0
	right := len(matrix[0]) - 1

	for left <= right && top <= bottom {

		//left to right
		for i := left; i <= right; i++ {
			res = append(res, matrix[top][i])
		}
		top++
		//top to bottom
		for i := top; i <= bottom; i++ {
			res = append(res, matrix[i][right])
		}
		right--
		//right to left
		if top <= bottom {
			for i := right; i >= left; i-- {
				res = append(res, matrix[bottom][i])
			}
			bottom--
		}
		//bottom to top
		if left <= right {
			for i := bottom; i >= top; i-- {
				res = append(res, matrix[i][left])
			}
			left++
		}

	}

	return res

}

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	res := printSpiralMatrix(matrix)
	fmt.Println(res)
}
