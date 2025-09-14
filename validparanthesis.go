package main

import "fmt"

func isValid(str string) bool {

	valid := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	stack := []rune{}

	for _, val := range str {
		if open, exist := valid[val]; exist {
			if len(stack) == 0 || stack[len(stack)-1] != open {
				return false
			}
			stack = stack[:len(stack)-1]

		} else {
			stack = append(stack, val)
		}
	}
	return len(stack) == 0

}

func main() {
	vals := "()[]{}"

	res := isValid(vals)
	fmt.Println(res)
}
