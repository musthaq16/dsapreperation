package main

import "fmt"

func fact(n int) int {

	count := 0

	for n > 0 {
		n = n / 5
		fmt.Println("nn", n)
		count = count + n
	}
	return count
}

func main() {
	val := 20
	res := fact(val)
	fmt.Println(res)
}
