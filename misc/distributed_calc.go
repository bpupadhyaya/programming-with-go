package main

import (
	"fmt"
	"runtime"
)

func sum(nums []int , chSum chan int) {
	sum := 0
	for _, value := range nums {
		sum += value
	}
	fmt.Println("func sum: ", sum)
	chSum <- sum
}

func product(nums []int , chProduct chan int) {
	product := 1
	for _, value := range nums {
		product *= value
	}
	fmt.Println("func product: ", product)
	chProduct <- product
}

func main() {
    fmt.Printf("Number of CPU cores: %d\n", runtime.NumCPU())
	chSum := make(chan int)
	myData := []int{3, 4, 5}
	go sum(myData, chSum)
	chProduct := make(chan int)
	go product(myData, chProduct)

	sum := <-chSum
	product := <-chProduct
	fmt.Println("main: sum: ", sum)
	fmt.Println("main: product: ", product)

}
