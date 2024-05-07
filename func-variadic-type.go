package main

import "fmt"

func sum(vals ...int) {
	fmt.Printf("%v %[1]T\n", vals)
}

func main() {
	sum(1, 2, 3, 4, 5)
	sum()
}
