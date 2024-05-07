package main

import "fmt"

func inc() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}

func main() {
	f1 := inc()
	f2 := inc()
	for i := 0; i < 5; i++ {
		fmt.Println("1", f1())
	}
	for i := 0; i < 3; i++ {
		fmt.Println("2", f2())
	}
	for i := 0; i < 5; i++ {
		fmt.Println("3", f1())
	}
}
