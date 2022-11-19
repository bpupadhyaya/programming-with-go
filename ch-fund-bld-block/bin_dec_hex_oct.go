package main

import (
    "fmt"
)

func main() {
    fmt.Printf("17 = decimal: %d\n",17)
    fmt.Printf("17 = hex: %x\n", 17)
    fmt.Printf("8 = octal: %o\n", 8)
    fmt.Printf("17 = binary: %b\n",17)
    fmt.Printf("a = %b binary; = %d decimal; =%x hex; =%o octal\n",'a', 'a', 'a', 'a')
    fmt.Printf("a = %b binary; = %d decimal; =%x hex; =%o octal\n",'A', 'A', 'A', 'A')
}