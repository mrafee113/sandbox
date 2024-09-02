package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "0001110010101010"
	fmt.Println(strings.TrimLeft(s, "0"))
	fmt.Println(strings.TrimPrefix(s, "0"))
}
