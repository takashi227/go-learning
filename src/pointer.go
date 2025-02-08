package main

import "fmt"

// ポインタ

func Double(n int) {
	n = n * 2
}

func Double2(i *int) {
	*i = *i * 2
}

func Double3(s []int) {
	for i, v := range s {
		s[i] = v * 2
	}
}

func main() {
	var n int = 100
	fmt.Println(n)

	// &を使って変数nのアドレスを取得
	fmt.Println(&n)

	Double(n)

	fmt.Println(n)

	// ポインタ型の変数を宣言
	var p *int = &n
	fmt.Println(p)
	// ポインタ型の変数を参照
	fmt.Println(*p)

	// ポインタ型の変数を使って値を更新
	// *p = 300
	// fmt.Println(n)
	// n = 200
	// fmt.Println(*p)

	Double2(&n)
	fmt.Println(n)

	Double2(p)
	fmt.Println(*p)

	var sl []int = []int{1, 2, 3}
	fmt.Println(sl)

	Double3(sl)
	fmt.Println(sl)

}
