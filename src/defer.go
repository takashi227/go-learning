package main

import (
	"fmt"
	"os"
)

/*
 * defer
 * 関数の終了時に必ず実行される処理を登録する
 */
func deferTest() {
	defer println("deferTest end")
	println("deferTest start")
}

func runDefer() {
	defer println("runDefer end")
	defer println("1")
	defer println("2")
	defer println("3")
	println("runDefer start")
}

func main() {
	defer func() {
		fmt.Println("1")
		fmt.Println("2")
		fmt.Println("3")
	}()
	deferTest()
	runDefer()

	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	file.Write([]byte("Hello, World!"))
}
