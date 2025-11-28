package main

import "fmt"

func main() {
	// TODO: declare a type for Student (with first and last name)
	type Student struct {
		FistName string
		LastName string
	}
	// TODO: declare a type for Class (consisting of multiple students)
	type Class struct {
		Students []Student
	}
	// TODO: declare a map of modules being attended by multiple classes
	classA := Class{Students: []Student{
		{"John", "Doe"},
		{"Jane", "Doe"}}}
	modules := map[int]Class{
		104: classA,
	}
	// TODO: output everything using fmt.Println()
	fmt.Println(modules)
}
