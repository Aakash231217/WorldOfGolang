package main

import "fmt"

func divide(a, b int) int {
	return a / b
}

func main() {
	fmt.Println("Error handling in functions")
	ans := divide(10, 4)
	fmt.Println("Division of two numbers is", ans)
}
