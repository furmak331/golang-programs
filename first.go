package main

import "fmt"

func main() {
    // Print a simple message
    fmt.Println("Hello, World!")

    // Variables in Go
    var name string = "Furqan"
    age := 22 // Short variable declaration

    fmt.Println("Name:", name)
    fmt.Println("Age:", age)

    // If-else statement
    if age > 18 {
        fmt.Println(name, "is an adult.")
    } else {
        fmt.Println(name, "is not an adult.")
    }

    // A simple for loop
    for i := 1; i <= 5; i++ {
        fmt.Println("Loop no. :", i)
    }
}
