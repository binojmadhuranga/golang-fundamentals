package main

import "fmt"

func main() {
	// create slice
	numbers := []int{10, 20, 30}

	// print slice
	fmt.Println("Slice:", numbers)

	// access element
	fmt.Println("First element:", numbers[0])

	numbers[1] = 99
	fmt.Println("After update:", numbers)

	numbers = append(numbers, 40)
	fmt.Println("After append:", numbers)

	fmt.Println("Length:", len(numbers))
	fmt.Println("Capacity:", cap(numbers))
}