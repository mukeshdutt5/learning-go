package main

import (
	"fmt"
)

func main() {

	// Array with declaration and initialization
	nums := []int{10, 11, 12, 13}

	// when we use foreach first item is index and second item is value of array
	for i, val := range nums {

		fmt.Printf("Index is : %d, Value is : %d \n", i, val)
	}

	// Map data
	lang := map[string]string{"a": "Apple", "b": "Ball"}

	// When we use key
	for key, value := range lang {
		fmt.Printf("Key is : %s, value is : %s\n", key, value)
	}

}
