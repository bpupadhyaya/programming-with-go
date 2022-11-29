package main

import (
	"fmt"
)

func Filter[T any](myData []T, myFunc func(T) bool) []T {
	var holder []T
	for _, currentData := range myData {
		if myFunc(currentData) {
			holder = append(holder, currentData)
		}
	}
	return holder
}

func lenLessThanTen(myData string) bool {
	return len(myData) < 10
}

func main() {
	myBag := []string{"apple", "grape", "banana", "watermelon", "orange"}
	filteredMyBag := Filter(myBag, lenLessThanTen)
	fmt.Println(filteredMyBag)
}
