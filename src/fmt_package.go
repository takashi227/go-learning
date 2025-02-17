package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("表示")

	fmt.Print("改行しない")
	fmt.Println("改行する")

	// カンマ区切り（半角スペースで区切られて表示される）
	fmt.Println("1", "2", "3")
	fmt.Println("A", "B", "C")

	// Printf フォーマット指定子
	fmt.Printf("%s\n", "文字列")
	fmt.Printf("%#v\n", "文字列")

	// Sprintf 文字列を返す
	s := fmt.Sprint("文字列1")
	s2 := fmt.Sprintf("%v\n", "文字列2")
	s3 := fmt.Sprintln("文字列3")

	fmt.Println(s)
	fmt.Println(s2)
	fmt.Println(s3)

	// Fprintf 書き込み先を指定
	fmt.Fprint(os.Stdout, "標準出力")
	fmt.Fprintf(os.Stdout, "標準出力2")
	fmt.Fprintln(os.Stdout, "標準出力3")

	f, _ := os.Create("test1.txt")
	defer f.Close()

	fmt.Fprint(f, "ファイル出力")
}
