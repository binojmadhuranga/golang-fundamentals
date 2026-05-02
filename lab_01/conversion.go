package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println("=== Numeric Conversions ===")

	var intVal int = 42

	var floatVal float64 = float64(intVal)
	fmt.Println("int to float64:", floatVal)

	var floatNum float64 = 99.99
	var intFromFloat int = int(floatNum)
	fmt.Println("float64 to int:", intFromFloat)

	var bigInt int64 = int64(intVal)
	fmt.Println("int to int64:", bigInt)

	fmt.Println("\n=== String Conversions ===")

	strFromInt := strconv.Itoa(intVal)
	fmt.Println("int to string:", strFromInt)

	str := "123"
	intFromStr, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Error converting string to int:", err)
	} else {
		fmt.Println("string to int:", intFromStr)
	}

	floatStr := strconv.FormatFloat(floatNum, 'f', 2, 64)
	fmt.Println("float to string:", floatStr)

	floatFromStr, err := strconv.ParseFloat("45.67", 64)
	if err != nil {
		fmt.Println("Error converting string to float:", err)
	} else {
		fmt.Println("string to float:", floatFromStr)
	}

	fmt.Println("\n=== Boolean Conversions ===")

	boolStr := "true"
	boolVal, err := strconv.ParseBool(boolStr)
	if err != nil {
		fmt.Println("Error converting string to bool:", err)
	} else {
		fmt.Println("string to bool:", boolVal)
	}

	boolToStr := strconv.FormatBool(true)
	fmt.Println("bool to string:", boolToStr)

	fmt.Println("\n=== Mixed Type Conversions ===")

	a := 10
	b := 3.5

	result := float64(a) + b
	fmt.Println("int + float64:", result)

	fmt.Println("\n=== Byte & Rune Conversions ===")

	var char rune = 'A'
	var ascii int = int(char)
	fmt.Println("rune to int (ASCII):", ascii)

	var backToChar rune = rune(ascii)
	fmt.Println("int to rune:", string(backToChar))

	fmt.Println("\n=== Lossy Conversion ===")

	var largeFloat float64 = 12345.6789
	var smallInt int = int(largeFloat)
	fmt.Println("float to int (lossy):", smallInt)
}