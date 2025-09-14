package main

import "fmt"

func nthstair(n int) int {

	if n <= 2 {
		return n
	}
	return nthstair(n-1) + nthstair(n-2)

}
func main() {
	val := 4

	if val <= 2 {
		fmt.Println(val)
		return
	}

	a, b := 1, 2

	for i := 3; i <= val; i++ {
		a, b = b, a+b
	}
	fmt.Println(b)

	res := nthstair(val)
	fmt.Println(res)
}
