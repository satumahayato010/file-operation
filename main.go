package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// ファイルの生成
	file, err := os.Create("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// ファイルの存在確認
	_, err1 := os.Stat("test.txt")
	if os.IsNotExist(err1) {
		fmt.Println("no such file")
	}

	// ファイルを開く
	file1, err2 := os.Open("test.txt")
	if err != nil {
		log.Fatal(err2)
	}
	defer file1.Close()

	// ファイル削除
	err3 := os.Remove("test.txt")
	if err != nil {
		log.Fatal(err3)
	}

	// ファイル名変更
	err4 := os.Rename("test.txt", "test1.txt")
	if err != nil {
		log.Fatal(err4)
	}

	// ファイルの読み込み
	file2, err5 := os.Open("test.txt")
	if err != nil {
		log.Fatal(err5)
	}
	defer file2.Close()

	b := make([]byte, 64)
	n, err := file2.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b[:n]))

	//ファイルの書き込み
	file3, err := os.OpenFile("test.txt", os.O_RDWR|os.O_APPEND, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer file3.Close()
	data := []byte("Hello World!!!")
	_, err6 := file3.Write(data)
	if err != io.EOF && err != nil {
		log.Fatal(err6)
	}

	// ioutilでのファイル読み込み
	b, err7 := ioutil.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err7)
	}
	fmt.Println(string(b))

	// ioutilでのファイル書き込み
	data2 := []byte("Hello World!!!")
	err8 := ioutil.WriteFile("test.txt", data2, 0755)
	if err != nil {
		log.Fatal(err8)
	}
}
