package main

import (
	"fmt"
	"time"
)

func myFun1() {
	for i := 0; i < 20; i++ {
		fmt.Printf("1 ")
	}
	fmt.Println("F1 ends")
}

func myFun2() {
	for i := 0; i < 20; i++ {
		fmt.Printf("2 ")
	}
	fmt.Println("F2 ends")
}

func main() {
	go myFun1()
	go myFun2()
	time.Sleep(time.Second)
	fmt.Println("Main ends")
}
