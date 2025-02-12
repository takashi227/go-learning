package main

import (
	. "fmt" // ドットをつけるとパッケージ名を省略できる
	f "fmt" // "fmt"をfとしてパッケージ名を変更
	// "foo"
)

func appName() string {
	const AppName = "GoApp"
	var Version string = "1.0"
	return AppName + " " + Version
}

func Do(s string) (b string) {
	// var b string = s // 再定義になっているのでエラーになる
	b = s
	{
		b := "BBB"
		f.Println(b)
	}
	f.Println(b)
	return b
}

func main() {
	f.Println("Hello, World!")
	Println("Hello, World! 2")

	f.Println(appName())
	// f.Println(AppName)
	// f.Println(Version)

	f.Println(Do("AAA"))
}
