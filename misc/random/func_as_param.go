package main

import (
	"fmt"
)

func main() {
	calculate(product, 3, 5)
	calculate(sum, 5, 6)
}

func calculate(f func(int, int) int, a, b int) {
	fmt.Println(f(a, b))
}

func product(a, b int) int {
	return a * b
}

func sum(a, b int) int {
	return a + b
}
