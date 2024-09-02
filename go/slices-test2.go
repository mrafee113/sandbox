package main

import "fmt"

func main() {
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("arr:", arr)
	s1 := arr[2:8]
	fmt.Println("s1:", s1)
	s1 = append(s1, s1[len(s1)-5:len(s1)-1]...)
	fmt.Println("s1:", s1)
	// ---------
	var runes []rune
	runes = append(runes, []rune("hello, world!")...)
	fmt.Println(runes)
}
