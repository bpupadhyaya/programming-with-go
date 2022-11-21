package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	ioReader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	name, _ := ioReader.ReadString('\n')
	fmt.Println("Hi ", name)

}
