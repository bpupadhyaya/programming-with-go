package main

import (
	"fmt"
)

type Stack[T comparable] struct {
	data []T
}

func (myStack *Stack[T]) Push(myData T) {
	myStack.data = append(myStack.data, myData)
}

func (myStack *Stack[T]) Pop() (T, bool) {
	if len(myStack.data) == 0 {
		var none T
		return none, false
	}
	top := myStack.data[len(myStack.data)-1]
	myStack.data = myStack.data[:len(myStack.data)-1]
	return top, true
}

func (myStack Stack[T]) has(myData T) bool {
	for _, currentData := range myStack.data {
		if currentData == myData {
			return true
		}
	}
	return false
}

func main() {
	var myStack1 Stack[string]
	myStack1.Push("Ken")
	myStack1.Push("Ritchie")
	data, status := myStack1.Pop()
	fmt.Println(data, status)

	var myStack2 Stack[int]
	myStack2.Push(5)
	myStack2.Push(10)
	data2, status2 := myStack2.Pop()
	fmt.Println(data2, status2)

	fmt.Println(myStack1.has("Ken"))
	fmt.Println(myStack1.has("bhim"))
	fmt.Println(myStack2.has(5))

}
