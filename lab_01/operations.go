package main

import "fmt"

func main() {

    // Variables
    a := 10
    b := 3

    // 🔹 Arithmetic
    fmt.Println("Arithmetic Operations:")
    fmt.Println("Addition:", a+b)
    fmt.Println("Subtraction:", a-b)
    fmt.Println("Multiplication:", a*b)
    fmt.Println("Division:", a/b)
    fmt.Println("Modulus:", a%b)

    // 🔹 Relational
    fmt.Println("\nRelational Operations:")
    fmt.Println("a == b:", a == b)
    fmt.Println("a != b:", a != b)
    fmt.Println("a > b:", a > b)
    fmt.Println("a < b:", a < b)
    fmt.Println("a >= b:", a >= b)
    fmt.Println("a <= b:", a <= b)

    // 🔹 Logical
    x := true
    y := false

    fmt.Println("\nLogical Operations:")
    fmt.Println("x && y:", x && y)
    fmt.Println("x || y:", x || y)
    fmt.Println("!x:", !x)

   
}