package main

import "fmt"

func main() {
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s1 := arr[2:8]
	fmt.Println("arr:", arr, "s1:", s1, "len:", len(s1), "cap:", cap(s1))
	s1[0] = 100
	fmt.Println("arr:", arr)
	s1 = append(s1, 101, 102, 103, 104, 105, 106, 107)
	fmt.Println("arr:", arr, "s1", s1)
	s1[0] = 200
	s1[len(s1)-1] = 201
	fmt.Println("arr:", arr, "s1", s1)
}
