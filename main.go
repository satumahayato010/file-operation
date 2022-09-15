package main

import (
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("test.txt", os.O_RDWR|os.O_APPEND, 0775)
	if err != nil {
		log.Fatal(err)
	}
	w := []byte("write test\n")
	_, err = file.Write(w)
	if err != io.EOF && err != nil {
		log.Fatal(err)
	}
}
