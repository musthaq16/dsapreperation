package main

import "fmt"

func setMatrixZeros(matrix [][]int) {

	rowlen := len(matrix)
	collen := len(matrix[0])

	rowarr := make([]bool, rowlen)
	colarr := make([]bool, collen)

	for i := 0; i < rowlen; i++ {
		for j := 0; j < collen; j++ {
			if matrix[i][j] == 0 {
				colarr[j] = true
				rowarr[i] = true
			}
		}
	}

	for i := 0; i < rowlen; i++ {
		for j := 0; j < collen; j++ {
			if colarr[j] || rowarr[i] {
				matrix[i][j] = 0
			}
		}
	}

	fmt.Println(matrix)

}

func main() {
	matrix := [][]int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	}

	setMatrixZeros(matrix)
}
