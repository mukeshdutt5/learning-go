package main

import (
	"fmt"
)

var data map[int]string

func main() {

	data = make(map[int]string)
	data[2] = "Hello"
	data[3] = "Hellor"

	delete(data, 2)

	fmt.Println(data)

	programming := map[int]string{1: "C#", 2: "golang", 3: "NodeJs"}

	fmt.Println(programming)
}
