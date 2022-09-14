package main

import (
	"log"
	"os"
)

func main() {
	err := os.Rename("test.txt", "test1.txt")
	if err != nil {
		log.Fatal(err)
	}
}
