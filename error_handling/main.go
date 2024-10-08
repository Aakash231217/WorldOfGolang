package main

import "fmt"

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("Denominator must not be zero")
	}
	return a / b, nil
}

func main() {
	fmt.Println("Error handling in functions")
	ans, _ := divide(10, 4)
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error Handling")
	}
	fmt.Println("Division of two numbers is", ans, result)
}
