package main

import (
	"fmt"
)

func main() {
	a := 4
	if a < 3 {
		fmt.Println("a is less than 3")
	} else if a < 5 {
		fmt.Println("a is less than 5 and greater than 3")
	} else {
		fmt.Println("a is greater than 5")
	}
}
