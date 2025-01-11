package main

import "fmt"

func main() {
	var i int = 100
	var i2 int64 = 200

	// 型が異なるため、エラーが発生する
	// fmt.Println(i + i2)

	// 型を調べる
	fmt.Printf("i = %T\n", i)
	fmt.Printf("i2 = %T\n", i2)

	// 型を変換する
	fmt.Printf("cast i2 = %T\n", int32(i2))
}
