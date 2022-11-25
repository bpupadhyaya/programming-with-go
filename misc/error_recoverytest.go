package main

import (
	"fmt"
	"log"
)

func myFunc() error {
	defer fmt.Println("Line 1")
	fmt.Println("Line 2")
	return fmt.Errorf("function terminated abruptly!")
	fmt.Println("Line 3")
	return nil
}

func main() {
	err := myFunc()
	if err != nil {
		log.Fatal(err)
	}
}
