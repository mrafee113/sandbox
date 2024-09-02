package main

import "fmt"

func main() {
	var s []string
	fmt.Println(s, len(s), cap(s), s == nil)
	s = append(s, "a")
	fmt.Println(s, len(s), cap(s), s == nil)
	s = nil
	fmt.Println(s, len(s), cap(s), s == nil)
	s1 := []string{}
	fmt.Println(s1, len(s1), cap(s1), s1 == nil)
}
