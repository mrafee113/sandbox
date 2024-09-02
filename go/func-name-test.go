package main

import "fmt"

func square(n int) int { return n * n }

func main() {
	var f func(int) int = square
	fmt.Printf("%v, %[1]T\n", f)

	var mymap = make(map[string]func(int) int)
	fmt.Println(mymap)
	mymap["square"] = square
	fmt.Println(mymap)
}
