package main

import (
	"fmt"
)

func main() {

	fmt.Println(factorial(7))

}

func factorial(number int) int {

	if number == 0 {
		return 1
	}

	return number * factorial(number-1)
}
