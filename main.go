package main

import (
	"fmt"
	"mylearning/myutil"
)

func main() {
	fmt.Println("Learn Go language by hello world")
	myutil.PrintMessage("HelloWorld")

	var name string = "aakash"
	fmt.Println(name)
	var version = "version1.0"
	fmt.Println(version)
	const pi = 6.402
	fmt.Println(pi)

	person := "aakash singh"
	fmt.Println(person)

	var Public = "data is important"
	var private = "data is private"

	fmt.Println(Public)
	fmt.Println(private)
	fmt.Printf("height is %.2f\n", pi)
}
