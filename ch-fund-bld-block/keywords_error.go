package main

import (
    "fmt"
)

func main() {
    var goto int // syntax error: unexpected goto, expecting name
    goto = 5 //  syntax error: unexpected =, expecting name
    fmt.Println("goto: ", goto) // syntax error: unexpected goto, expecting expression
}