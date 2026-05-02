package main

import (
	"fmt"
)

func main() {

	
	var a int = 10
	var b int8 = 127
	var c int16 = 32000
	var d int32 = 2000000000
	var e int64 = 9000000000000

	
	var u uint = 20
	var u8 uint8 = 255
	var u16 uint16 = 65000
	var u32 uint32 = 4000000000
	var u64 uint64 = 18000000000000000000

	
	var f1 float32 = 3.14
	var f2 float64 = 3.1415926535

	
	var isGoAwesome bool = true

	
	var name string = "Binoj"

	
	x := 100
	y := "Hello Go"


	var numbers [3]int = [3]int{1, 2, 3}

	
	slice := []int{10, 20, 30}

	
	student := map[string]int{
		"math":    90,
		"science": 85,
	}

	
	p := &a

	
	type Person struct {
		name string
		age  int
	}


	person1 := Person{name: "Binoj", age: 25}

	var zeroInt int
	var zeroString string
	var zeroBool bool

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



   //Example use of short variable declaration

	name2 := "Binoj2"
	age2 := 25

	fmt.Println("Name:", name2)
	fmt.Println("Age:", age2)




}