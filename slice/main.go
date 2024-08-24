package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	numbers = append(numbers, 3, 10, 12, 14, 15, 16, 19)
	fmt.Println("Number:", numbers)
	fmt.Printf("Number has a data type:%T\n", numbers)
	fmt.Println("Length", len(numbers))

	name := []string{}
	fmt.Println("name:", name)

	number := make([]int, 3, 5)
	number = append(number, 4)
	number = append(number, 98)
	number = append(number, 6)
	fmt.Println("Slice", number)
	fmt.Println("Length", len(number))
	fmt.Println("Capacity", cap(number))

	//var numbers = []string // you cant define slice this

	stock := make([]int, 0)
	fmt.Println("Slice", stock)
	fmt.Println("")

}
