package main

import "fmt"

func rmdup(strings []string) []string {
	counter := 0
	for i := 0; i < len(strings)-1; i++ {
		if strings[i] == strings[i+1] {
			copy(strings[i:], strings[i+1:])
			counter++
		}
	}
	return strings[:counter+1]
}

func main() {
	s := []string{"a", "b", "c", "c", "d", "e", "e", "f", "f"}
	fmt.Println(s)
	s = rmdup(s)
	fmt.Println(s)
}
