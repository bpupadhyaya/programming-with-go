package main
import (
    "fmt"
)

func myFurtherFunc() {
    p := recover()
    err1, ok := p.(error)
    if ok {
        fmt.Println(err1.Error())
    }
    fmt.Println("After recovery line")
}

func myProblematic() {
    defer myFurtherFunc()
    err := fmt.Errorf("Validatoin error")
    panic(err)

}

func main() {
    myProblematic()
    fmt.Println("After call to problematic")
}
