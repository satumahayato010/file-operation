package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	read := strings.NewReader("Hello World!!!")
	b, err := ioutil.ReadAll(read)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
}
