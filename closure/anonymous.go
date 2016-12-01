package main

import "fmt"

func printMessage(message string) {
	fmt.Println(message)
}

func getPrintMessage() func(string) {

	return func(message string) {
		fmt.Println(message)
	}
}

func getResult() func() {

	// When you call fuction as non return type
	fmt.Println("Step1 here")

	// When you call this fuction as return type function
	return func() {
		fmt.Println("Message from closure function")
	}
}

func main1() {

	printMessage("Hello named fuction")

	// Function declare and called at same place
	func(message string) {
		fmt.Println(message)
	}("delcare and called")

	printfunc := getPrintMessage()
	printfunc("Hello anonymous function")
}
