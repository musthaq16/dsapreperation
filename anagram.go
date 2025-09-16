package main

import "fmt"

func validAnagram(s, t string) bool {

	seen := make(map[rune]int)

	if len(s) != len(t) {
		return false
	}

	for _, val := range s {
		seen[val]++
	}

	for _, val := range t {
		seen[val]--
	}

	for _, val := range seen {
		if val == 0 {
			continue
		} else {
			return false
		}
	}

	return true
}

func main() {
	s := "anagramn"
	t := "nagaraml"

	valid := validAnagram(s, t)
	fmt.Println(valid)
}
