package main

import "fmt"

func main() {
    // No need to set value in explicit type at time of declaration
    // Type will be set based on the type
    var anInteger int = 1
    var anChar int32 = 'a'
    var anString string = "I am string"

    fmt.Println(anInteger)
    fmt.Println(anChar)
    fmt.Println(anString)

    fmt.Println()

    fmt.Printf("%T\n", anInteger)
    fmt.Printf("%T\n", anChar)
    fmt.Printf("%T\n", anString)

}
