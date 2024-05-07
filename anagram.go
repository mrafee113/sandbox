package main

import "fmt"

func anagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	c1 := make(map[rune]int)
	c2 := make(map[rune]int)
	var chars []rune
	for _, char := range s1 {
		c1[char]++
		chars = append(chars, char)
	}
	for _, char := range s2 {
		c2[char]++
		chars = append(chars, char)
	}
	for _, char := range chars {
		if c1[char] != c2[char] {
			return false
		}
	}
	return true
}

func main() {
	list := [][]string{
		[]string{"debit card", "bad credit"},
		[]string{"dormitory", "dirtyroom"},
		[]string{"abcd", "abcde"},
		[]string{"abcd", "dcba"},
	}
	for _, str := range list {
		s1, s2 := str[0], str[1]
		fmt.Println(s1, s2, anagram(s1, s2))
	}
}
