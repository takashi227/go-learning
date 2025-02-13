package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// fmt.Println("START")
	// os.Exit(1)
	// fmt.Println("END")

	// _, err := os.Open("A.txt")
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// ファイル操作
	// f, err := os.Open("test.txt")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// defer f.Close()

	// ファイル作成
	// f, _ := os.Create("test2.txt")
	// // 書き込み
	// f.Write([]byte("Hello, World\n"))
	// f.WriteAt([]byte("Golang\n"), 6)
	// f.Seek(0, os.SEEK_END)
	// f.WriteString("Yaah")

	// ファイルの読み込み
	// f, err := os.Open("test2.txt")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// defer f.Close()

	// // バイト列を読み込む
	// b := make([]byte, 128)
	// n, _ := f.Read(b)
	// fmt.Println(n)
	// fmt.Println(string(b))

	// b2 := make([]byte, 128)
	// nn, _ := f.ReadAt(b2, 10)
	// fmt.Println(nn)
	// fmt.Println(string(b2))

	// OpenFile
	f, err := os.OpenFile("test2.txt", os.O_RDONLY, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	// バイト列を読み込む
	b := make([]byte, 128)
	n, _ := f.Read(b)
	fmt.Println(n)
	fmt.Println(string(b))

}
