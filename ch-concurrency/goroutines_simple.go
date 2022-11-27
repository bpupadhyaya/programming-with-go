package main

import (
	"fmt"
	"time"
)

func myFun1(myParam string) {
	fmt.Println("myFun1: ", myParam)
}

func myFun2(myParam string) {
	fmt.Println("myFun2: ", myParam)
}

func main() {
	fmt.Println("main: start")
	go myFun1("hi")
	go myFun2("hi")
	time.Sleep(2 * time.Second)
	fmt.Println("main: end")
}
