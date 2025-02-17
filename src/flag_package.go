package main

import (
	"flag"
	"fmt"
)

func main() {
	// オプションの値を格納する変数の定義
	var (
		max int
		msg string
		x   bool
	)

	// オプションの解析
	// IntVar 整数のオプション
	flag.IntVar(&max, "n", 256, "処理数の最大値")
	// StringVar 文字列のオプション
	flag.StringVar(&msg, "m", "Hello, World", "処理メッセージ")
	// BoolVar bool値のオプション コマンドラインにオプションがあればtrue
	flag.BoolVar(&x, "x", false, "拡張オプション")

	// コマンドラインをパース
	flag.Parse()

	fmt.Println("処理数の最大値:", max)
	fmt.Println("処理メッセージ:", msg)
	fmt.Println("拡張オプション:", x)

	// コマンドラインを処理するサンプル
	// go run flag_package.go -n 512 -m "Hello, Golang" -x
}
