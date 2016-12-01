package main

import "fmt"

const (
	a = iota
	b
	c
	d
)

func main() {

	var i int
	i = 5

	s := make([]int, i)

	s[0] = 11
	s[1] = 11
	s[2] = 11
	s[3] = 11
	s[4] = 11

	// Getting length of slice or array
	fmt.Println(len(s))

	// Appending elements into slice or array
	s = append(s, 10)

	s2 := make([]int, len(s))

	// Coping one array to another array
	copy(s2, s)

	fmt.Println(s)
	fmt.Println(s2)
}
