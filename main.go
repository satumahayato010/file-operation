package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Stat("test.txt")
	if os.IsNotExist(err) {
		fmt.Println("no such file")
	}
}
