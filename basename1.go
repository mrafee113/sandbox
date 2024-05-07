package main

import (
	"fmt"
	"strings"
)

func basename(s string) string {
	s = s[strings.LastIndex(s, "/")+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

func main() {
	for ndx, s := range []string{"a/b/c.go", "c.d.go", "abc"} {
		sbased := basename(s)
		fmt.Println(ndx, s, sbased)
	}
}
