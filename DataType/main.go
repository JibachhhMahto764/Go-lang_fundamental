package main

import (
        "fmt"
        "strings"
)

func main() {
        x := func(s string) string {
                return strings.ToUpper(s)
        }
        fmt.Printf("%T \n", x)
        fmt.Println(x("Joe"))
}