package main

import "fmt"

func test(x int) (z int) {
	z = x
	return
}

func main() {
	v := test(2)
	fmt.Println(v)
}
