package main

import (
	"fmt"
)

func myFun1(ch chan string) {
	ch <- "hello from myFun1"
}

func main() {
	myChan := make(chan string)
	go myFun1(myChan)
	fmt.Println(<-myChan)
}
