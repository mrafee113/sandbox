package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	var w io.Writer
	fmt.Printf("%T %[1]v\n", w)
	w = os.Stdout
	fmt.Printf("%T %[1]v\n", w)
}
