package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	res1, err1 := http.Get("https://www.google.com/")
	if err1 != nil {
		log.Fatal(err1)
	}
	res2, err2 := http.Get("https://www.facebook.com/")
	if err2 != nil {
		log.Fatal(err2)
	}
	defer res2.Body.Close()
	defer res1.Body.Close()
	googBody, err := ioutil.ReadAll(res1.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Google content")
	fmt.Printf("Goog size: %d\n", len(googBody))

	fbBody, err := ioutil.ReadAll(res2.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Facebook content")
	fmt.Printf("Facebook size: %d\n", len(fbBody))
}
