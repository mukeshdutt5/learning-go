package main

import (
	"fmt"
)

func main() {

	data(10, 1021, 2121, 12121, 1212121, 212)
}

// Example of variadic function in golang
func data(nums ...int) {

	var sum int
	for _, value := range nums {
		sum += value
	}

	fmt.Printf("Total sum passed by parameter is : %d\n", sum)
}
