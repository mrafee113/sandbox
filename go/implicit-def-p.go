package main

import "fmt"

func main() {
	type Track struct {
		Title, Artist string
	}

	var tracks = []Track{
		{"Go", "Delilah"},
		{"Ready", "Moby"},
	}
	fmt.Printf("%T %[1]v\n", tracks)
	fmt.Printf("%T %[1]v\n", tracks[0])
}
