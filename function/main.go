package main

import "fmt"

func simpleFunction() {
	fmt.Println("Simple function")
}

func add(a, b int) (result int) {
	result = a + b
	return result
}

func multiply(a, b int) (res int) {
	res = a * b
	return res
}

func main() {
	fmt.Println("We are learning functionss")
	simpleFunction()
	ans := add(3, 4)
	fmt.Println("Sum of two numbers", ans)
	result := multiply(2, 4)
	fmt.Println("Multiply of two numbers", result)
}
