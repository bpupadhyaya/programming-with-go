package main

import (
    "fmt"
)

func myFunct() {
    defer fmt.Println("First line")
    defer fmt.Println("Second line")
    fmt.Println("Third line")
    fmt.Println("Fourth line")
}

func main() {
    defer fmt.Println("main: first line")
    defer fmt.Println("main: second line")
    myFunct()
    fmt.Println("main: last line")
}