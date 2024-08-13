package main

import "fmt"

func main() {
	fmt.Println("We are learning array today")

	var name [5]string
	name[0] = "prince"
	name[2] = "aakash"

	fmt.Println("Name of person is:", name)

	var numbers = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Number is", numbers)

	fmt.Println("Length of Numbers array is:", len(numbers))
	fmt.Println("value of name at 2 index is", name[2])

	var price [5]int
	fmt.Println("Price is", price)

	var prices [5]bool
	fmt.Println("Price is", prices)

	var pricess [5]string
	fmt.Println("Price is", pricess)
	fmt.Println("Names are", name)
}
