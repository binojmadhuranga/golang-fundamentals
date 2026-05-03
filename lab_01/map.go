package main

import "fmt"

func main() {
	// map create කරනවා
	studentMarks := make(map[string]int)

	// values add කරනවා
	studentMarks["Binoj"] = 85
	studentMarks["Kamal"] = 75
	studentMarks["Nimal"] = 90

	// print map
	fmt.Println("All Marks:", studentMarks)

	// specific value access කරනවා
	fmt.Println("Binoj Marks:", studentMarks["Binoj"])

	// update value
	studentMarks["Binoj"] = 95

	// delete value
	delete(studentMarks, "Kamal")

	// loop through map
	for key, value := range studentMarks {
		fmt.Println(key, ":", value)
	}
}