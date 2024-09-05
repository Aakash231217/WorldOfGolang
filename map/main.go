package main

import "fmt"

func main() {
	//name->grade
	studentGrades := make(map[string]int)
	studentGrades["Aakash"] = 34
	studentGrades["Prince"] = 90
	studentGrades["Abhishek"] = 85

	fmt.Println("Marks of Aakash:", studentGrades["Aakash"])

	studentGrades["Aakash"] = 100
	fmt.Println("new marks of Aakash is", studentGrades["Aakash"])

	delete(studentGrades, "Prince")
	fmt.Println("Marks of Prince is", studentGrades["Prince"])

	grades, exists := studentGrades["Aakash"]
	fmt.Println("Grade of David is", grades)
	fmt.Println("David exists:", exists)

	for index, value := range studentGrades {
		fmt.Printf("Key is %s and marks is %d\n", index, value)
	}

	person := map[string]int{
		"Alice":   90,
		"Bob":     86,
		"Charlie": 95,
	}

	for index, value := range person {
		fmt.Printf("----key is %s and marks is %d\n", index, value)
	}

}
