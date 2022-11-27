package main

import (
	"fmt"
)

func myFun1(ch chan string) {
	ch <- "1"
	ch <- "3"
}

func myFun2(ch chan string) {
	ch <- "2"
	ch <- "4"
}

func main() {
	myChan1 := make(chan string)
	myChan2 := make(chan string)
	go myFun1(myChan1)
	go myFun2(myChan2)
	fmt.Println(<-myChan1)
	fmt.Println(<-myChan2)
	fmt.Println(<-myChan1)
	fmt.Println(<-myChan2)
}
