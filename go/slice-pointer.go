package main

import "fmt"

func test(s *[]int, v int) {
	s1 := *s
	s1 = append(s1, v)
	fmt.Println(s1, *s)
	*s = s1
}

func main() {
	v := []int{1, 2, 3, 4, 5}
	fmt.Println(v)
	test(&v, 100)
	fmt.Println(v)
}
