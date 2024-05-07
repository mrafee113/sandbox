package main

import "fmt"

func main() {
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s1 := arr[2:8]
	s2 := append(s1, 100, 200)
	fmt.Println("arr:", arr)
	fmt.Println("s1:", s1, "len(s1):", len(s1), "cap(s1):", cap(s1))
	fmt.Println("s2:", s2, "len(s2):", len(s2), "cap(s2):", cap(s2))
}
