package main

import (
	"crypto/sha256"
	"fmt"
)

func bitdiff(d1, d2 [32]byte) uint8 {
	fmt.Println(d1)
	fmt.Println(d2)
	var counter uint8 = 0
	for ndx := 0; ndx < 32; ndx++ {
		fmt.Println(ndx, d1[ndx], d2[ndx])
		if d1[ndx] != d2[ndx] {
			counter++
		}
	}
	return counter
}

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Println("x", "X")
	fmt.Println(bitdiff(c1, c2))
}
