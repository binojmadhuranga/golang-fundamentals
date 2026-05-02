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

    // 🔹 Assignment
    c := 5
    fmt.Println("\nAssignment Operations:")
    c += 2
    fmt.Println("c += 2:", c)
    c -= 1
    fmt.Println("c -= 1:", c)
    c *= 3
    fmt.Println("c *= 3:", c)
    c /= 2
    fmt.Println("c /= 2:", c)

}