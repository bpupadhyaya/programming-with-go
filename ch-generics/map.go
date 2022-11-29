package main

import (
	"fmt"
)

func Map[T1, T2 any](myData []T1, myFunc func(T1) T2) []T2 {
	holder := make([]T2, len(myData))
	for i, currentData := range myData {
		holder[i] = myFunc(currentData)
	}
	return holder
}

func lenFunc(myString string) int {
	return len(myString)
}

func main() {
	fruits := []string{"apple", "grape", "banana"}
	strLengths := Map(fruits, lenFunc)
	fmt.Println(strLengths)

}
