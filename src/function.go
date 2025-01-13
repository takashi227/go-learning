package main

import "fmt"

// func Plus(x int, y int) int {
// 	return x + y
// }

// 引数の型が同じ場合、省略可能
func Plus(x, y int) int {
	return x + y
}

func Div(x, y int) (int, int) {
	q := x / y
	r := x % y
	return q, r
}

func Double(price int) (result int) {
	result = price * 2
	return
}

func NoReturn() {
	fmt.Println("No return")
}

// 関数を返す関数
func ReturnFunc() func() {
	return func() {
		fmt.Println("I'm a function")
	}
}

func main() {
	var i int = Plus(1, 2)
	fmt.Println(i)

	var q, r int = Div(9, 2)
	fmt.Println(q, r)

	// 戻り値の一部を無視する
	q, _ = Div(20, 3)
	fmt.Println(q)

	fmt.Println(Double(100))

	NoReturn()

	// 無名関数
	f := func(x, y int) int {
		return x + y
	}
	i2 := f(1, 2)
	fmt.Println(i2)

	// 即時関数
	i3 := func(x, y int) int {
		return x + y
	}(1, 2)
	fmt.Println(i3)

	f2 := ReturnFunc()
	f2()
}
