package main

import "fmt"

func longestCommonPrefix(strs []string) string {

	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]

	for i := 1; i < len(strs); i++ {
		currStr := strs[i]

		for len(prefix) > 0 && (len(prefix) > len(currStr) || prefix != currStr[:len(prefix)]) {
			prefix = prefix[:len(prefix)-1]
		}

		if prefix == "" {
			return ""
		}
	}

	return prefix

}

func main() {

	vals := []string{"flower", "flow", "flight"}

	res := longestCommonPrefix(vals)

	fmt.Println(res)

}
