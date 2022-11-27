package main

import (
	"fmt"
	"time"
)

func mySleep(agent string, myDelay int) {
	for i := 0; i < myDelay; i++ {
		fmt.Println(agent, " sleeping")
		time.Sleep(1 * time.Second)
	}
}
func myFun1(ch chan string) {
	fmt.Println("myFun1 sending value 11 now...")
	ch <- "f1-1"
	//mySleep("myFun1: ", 2)
	fmt.Println("myFun1 sending value 12 now...")
	ch <- "f1-2"
	fmt.Println("myFun1 sending value 13 now...")
	ch <- "f1-3"
}

func myFun2(ch chan string) {
	fmt.Println("myFun2 sending value 21 now...")
	ch <- "f2-1"
	//mySleep("myFun2: ", 3)
	fmt.Println("myFun2 sending value 22 now...")
	ch <- "f2-2"
	fmt.Println("myFun2 SENT value 22 ")
	fmt.Println("myFun2 sending value 23 now...")
	ch <- "f2-3"
}

func main() {
	myChan1 := make(chan string)
	myChan2 := make(chan string)
	go myFun1(myChan1)
	go myFun2(myChan2)
	//mySleep("main: ", 3)
	fmt.Println("received11: ", <-myChan1)
	fmt.Println("received21: ", <-myChan2)
	//fmt.Println("received12: ", <-myChan1)
	fmt.Println("received22: ", <-myChan2)
	//mySleep("main: ", 3)
	fmt.Println("received13: ", <-myChan1)
	fmt.Println("received23: ", <-myChan2)
	fmt.Println("end of main func ")
}
