package main

import "fmt"

// Prints 122 to the console.
//
// This is the entry point of the application.
func main() {
	fmt.Println(soma(112, 10))
}

func soma(a int, b int) int {
	return a + b
}
