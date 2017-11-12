package main

import "fmt"

func sum(a, b int) int {
	return a + b
}

func main() {
	fmt.Println("monolith example")
	fmt.Printf("The sum of 3 and 4 is: %v\n", sum(3, 4))
}
