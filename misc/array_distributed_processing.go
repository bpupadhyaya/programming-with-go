package main

import (
	"fmt"
	"runtime"
	"time"
)

func calcSum(nums []int, chSum chan int) {
	total := 0
	for _, value := range nums {
		total += value
	}
	fmt.Println("calcSum: partial total: ", total)
	chSum <- total
}

func main() {
	chSum := make(chan int)
	grandTotal := 0
	myData := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	availableCores := runtime.NumCPU()
	fmt.Println("Number of CPU cores available: ", availableCores)
	chunkLength := len(myData) / availableCores // refine this logic to cover all cases later
	marker := 0
	var t1 int64 = time.Now().UnixNano()
	for i := 0; i < availableCores; i++ {
		go calcSum(myData[marker:marker+chunkLength], chSum)
		result := <-chSum
		grandTotal += result
		marker += chunkLength
	}
	var t2 int64 = time.Now().UnixNano()
	fmt.Println("Grand total from conc. proc.: ", grandTotal)

	grandTotal = 0
	var t3 int64 = time.Now().UnixNano()
	for _, value := range myData {
		grandTotal += value
	}
	var t4 int64 = time.Now().UnixNano()
	concTime := t2 - t1
	serTime := t4 - t3
	fmt.Println("Grand total from serial proc.: ", grandTotal)
	fmt.Printf("Concurrent time taken: %d nano seconds\n", concTime)
	fmt.Printf("Serial time taken: %d nano seconds\n", serTime)

}
