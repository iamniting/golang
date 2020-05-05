package main

import "fmt"

func main() {
    // Always need to set value in implicit type at time of declaration
    // Type will be set based on value
    anInteger := 1
    anChar := 'a'
    anString := "I am string"

    fmt.Println(anInteger)
    fmt.Println(anChar)
    fmt.Println(anString)

    fmt.Println()

    fmt.Printf("%T\n", anInteger)
    fmt.Printf("%T\n", anChar)
    fmt.Printf("%T\n", anString)
}
