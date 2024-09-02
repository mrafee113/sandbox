package main

import "fmt"

func comma(s string) string {
	end := 0
	if s[0] == '+' {
		s = s[1:]
	} else if s[0] == '-' {
		end = 1
	}
	start := len(s)
	for ndx := len(s) - 1; ndx >= 0; ndx-- {
		if s[ndx] == '.' {
			start = ndx
			break
		}
	}
	for ndx := start - 3; ndx > end; ndx -= 3 {
		s = s[:ndx] + "," + s[ndx:]
	}
	return s
}

func main() {
	for ndx, s := range []string{"12345", "+210987654321", "-123456.123456", "123456.14a12321312332"} {
		fmt.Println(ndx, s, comma(s))
	}
}
