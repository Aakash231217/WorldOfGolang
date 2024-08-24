package main

import "fmt"

func main() {
	x := 7
	if x == 10 {
		fmt.Println("x is greater than 5")
	} else {
		fmt.Println("x is not greater than 5")
	}
	z := 10
	if z > 10 {
		fmt.Println("z is greater than 5")
	} else {
		fmt.Println("z is smaller than 5")
	}

	y := 10
	if x > 5 && y > 5 {
		fmt.Println("Hey how are you")
	}
}
