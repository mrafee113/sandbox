package main

import "fmt"

func reverse(s *[]int) {
	var slice []int = *s
	for i, j := 0, len(*s)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

func main() {
	arr := [3]int{1, 2, 3}
	s1 := arr[:]
	fmt.Printf("%v %T\n", &s1, &s1)
	fmt.Println(s1[0])
	reverse(&s1)
	fmt.Println(s1)
}
