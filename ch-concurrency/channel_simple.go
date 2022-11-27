package main

import (
	"fmt"
)

func myFun(ch chan string) {
	ch <- "hi"
}

func main() {
	fmt.Println("main: begin")
	myChan := make(chan string)
	go myFun(myChan)
	fmt.Println(<-myChan)
	fmt.Println("main: end")
}
