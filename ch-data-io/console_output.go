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
	name, e := inputReader.ReadString('\n')
	if e != nil {
		log.Fatal(e)
	}
	fmt.Println("Hi ", name)
}
