package main

import (
	"fmt"
)

func main() {

	// 🔹 1. Integer types
	var a int = 10
	var b int8 = 127
	var c int16 = 32000
	var d int32 = 2000000000
	var e int64 = 9000000000000

	// 🔹 2. Unsigned integers
	var u uint = 20
	var u8 uint8 = 255
	var u16 uint16 = 65000
	var u32 uint32 = 4000000000
	var u64 uint64 = 18000000000000000000

	// 🔹 3. Floating point
	var f1 float32 = 3.14
	var f2 float64 = 3.1415926535

	// 🔹 4. Boolean
	var isGoAwesome bool = true

	// 🔹 5. String
	var name string = "Binoj"

	// 🔹 6. Short declaration (type inference)
	x := 100
	y := "Hello Go"

	// 🔹 7. Array
	var numbers [3]int = [3]int{1, 2, 3}

	// 🔹 8. Slice (dynamic array)
	slice := []int{10, 20, 30}

	// 🔹 9. Map (key-value)
	student := map[string]int{
		"math":    90,
		"science": 85,
	}

	// 🔹 10. Pointer
	p := &a

	// 🔹 11. Struct
	type Person struct {
		name string
		age  int
	}

	person1 := Person{name: "Binoj", age: 25}

	// 🔹 12. Zero values
	var zeroInt int
	var zeroString string
	var zeroBool bool

	// 🔹 OUTPUT

	fmt.Println("Integers:", a, b, c, d, e)
	fmt.Println("Unsigned:", u, u8, u16, u32, u64)
	fmt.Println("Floats:", f1, f2)
	fmt.Println("Boolean:", isGoAwesome)
	fmt.Println("String:", name)

	fmt.Println("Short declaration:", x, y)

	fmt.Println("Array:", numbers)
	fmt.Println("Slice:", slice)

	fmt.Println("Map:", student)

	fmt.Println("Pointer value:", *p)

	fmt.Println("Struct:", person1)

	fmt.Println("Zero values:", zeroInt, zeroString, zeroBool)
}