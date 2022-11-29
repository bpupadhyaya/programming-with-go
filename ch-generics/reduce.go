package main

import (
	"fmt"
)

func Reduce[T1, T2 any](myData []T1, initializer T2, myFunc func(T2, T1) T2) T2 {
	holder := initializer
	for _, currentData := range myData {
		holder = myFunc(holder, currentData)
	}
	return holder
}

func accumulate(acc int, myVal int) int {
	return acc + myVal
}

func main() {
	strLengths := []int{5, 5, 6} // from map.go
	myTotal := Reduce(strLengths, 0, accumulate)
	fmt.Println(myTotal)
}
