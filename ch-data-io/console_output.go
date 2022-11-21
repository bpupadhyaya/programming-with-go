package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Print("Please enter your name: ")
	name, error := inputReader.ReadString('\n')
	if error != nil {
		log.Fatal(error)
	}
	fmt.Println("Hi ", name)
}
