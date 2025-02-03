package main

import "fmt"

func main() {
	// スライスの宣言
	var sl []int
	fmt.Println(sl)

	// スライスの初期化
	var sl2 []int = []int{100, 200}
	fmt.Println(sl2)

	sl3 := []string{"A", "B"}
	fmt.Println(sl3)

	// make関数を使ったスライスの初期化
	sl4 := make([]int, 5)
	fmt.Println(sl4)

	// スライスの要素へのアクセス
	sl2[0] = 1000
	fmt.Println(sl2)

	// スライスの要素の取得
	sl5 := []int{1, 2, 3, 4, 5}
	fmt.Println(sl5[2])

	// スライスの要素の取得
	fmt.Println(sl5[2:4])
	fmt.Println(sl5[:2])
	fmt.Println(sl5[2:])
	fmt.Println(sl5[:])

	// 最後の要素を取得
	fmt.Println(sl5[len(sl5)-1])
	// 最初と最後以外の要素を取得
	fmt.Println(sl5[1 : len(sl5)-1])

	// append
	sl6 := []int{1, 2, 3}
	fmt.Println(sl6)
	sl6 = append(sl6, 4)
	fmt.Println(sl6)
	sl6 = append(sl6, 5, 6, 7)
	fmt.Println(sl6)

	// make
	sl7 := make([]int, 5)
	fmt.Println(sl7)

	// len
	fmt.Println(len(sl6))
	fmt.Println(len(sl7))

	// capacity
	fmt.Println(cap(sl6))
	fmt.Println(cap(sl7))

	sl8 := make([]int, 5, 10)
	fmt.Println(sl8)
	fmt.Println(len(sl8))
	fmt.Println(cap(sl8))
	sl8 = append(sl8, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(sl8)
	fmt.Println(len(sl8))
	fmt.Println(cap(sl8))
}
