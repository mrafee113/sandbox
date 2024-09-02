package main

import (
	"fmt"
	"os"
)

func main() {
	s := os.Args[1]
	sep := " "
	for ndx, arg := range os.Args[2:] {
		fmt.Println(ndx)
		s += sep + arg
	}
	fmt.Println(s)
}
