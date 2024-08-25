package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println("Number is:", i)
	}

	//	count := 0
	//	for {
	//		fmt.Println("Infinite loop")
	//		count++
	//	}

	numbers := []int{1, 3, 4, 5, 3, 5, 3}
	for index, value := range numbers {
		fmt.Printf("Index of numbers is %d, and value is %d\n", index, value)
	}

	data := "Hello World"
	for index, value := range data {
		fmt.Printf("Index of data is %d, and value is %c\n", index, value)
	}
}
