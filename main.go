package main

import (
	"io/ioutil"
	"log"
)

func main() {
	data := []byte("Hello World!!")
	err := ioutil.WriteFile("test.txt", data, 0755)
	if err != nil {
		log.Fatal(err)
	}
}
