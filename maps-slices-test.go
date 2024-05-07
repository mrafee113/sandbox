package main

import "fmt"

func k(list []string) string { return fmt.Sprintf("%q", list) }
func main() {
	//var m = make(map[string]int)
	strings := []string{"abc", "def", "ghi", "jkl", "mno"}
	fmt.Println(k(strings))
	fmt.Printf("%#v\n", k(strings))
}
