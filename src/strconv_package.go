package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 真偽値を文字列に変換する
	bt := true
	fmt.Printf("%T\n", strconv.FormatBool(bt))

	// 整数を文字列に変換する
	i := strconv.FormatInt(-100, 10)
	fmt.Printf("%v, %T\n", i, i)

	// 簡易的に変換できる
	i2 := strconv.Itoa(100)
	fmt.Printf("%v, %T\n", i2, i2)
}
