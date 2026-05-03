package main

import "fmt"

type Student struct {
	name  string
	age   int
	marks float64
}

func main() {
	var s1 Student

	s1.name = "Binoj"
	s1.age = 23
	s1.marks = 85.5

	fmt.Println("Name:", s1.name)
	fmt.Println("Age:", s1.age)
	fmt.Println("Marks:", s1.marks)
}