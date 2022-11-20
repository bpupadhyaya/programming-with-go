package main

import (
    "fmt"
)

func main() {
    var sampleVariable int
    sampleVariable = 678
    _sampleVariable := 8
    //1_sampleVariable := 8 // produces syntax error: "unexpected sample_variable at end of statement"
    fmt.Println("sampleVariable = ", sampleVariable)
    fmt.Println(" _sampleVariable = ", _sampleVariable)
}