package main

import "fmt"

func comma(s string) string {
	for ndx := len(s) - 3; ndx > 0; ndx -= 3 {
		s = s[:ndx] + "," + s[ndx:]
	}
	return s
}

func main() {
	for ndx, s := range []string{"12345", "210987654321"} {
		fmt.Println(ndx, s, comma(s))
	}
}
