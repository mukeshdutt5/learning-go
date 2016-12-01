package main

import "fmt"

func getOutput() (int, func()) {

	str := "string variable"

	// This function is closure fucntion because this is anonymous fucntion and it has access to get parent level variables
	foo := func() {
		fmt.Println(str)
	}

	foo()

	var a int
	a = 1

	return a, foo
}

func main() {
	// fmt.Println("He")

	a, f := getOutput()
	fmt.Println(a)
	f()
}
