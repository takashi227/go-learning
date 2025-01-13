package main

import "fmt"

/**
 * Go言語の配列の特徴
 * - 配列の要素数を変更できない
 */
func main() {
	var arr1 [3]int
	fmt.Println(arr1)
	fmt.Printf("arr1: %T\n", arr1)

	var arr2 [3]string = [3]string{"A", "B"}
	fmt.Println(arr2)
	fmt.Printf("arr2: %T\n", arr2)

	// 要素数を省略する
	var arr3 = [...]string{"X", "Y", "Z"}
	fmt.Println(arr3)
	fmt.Printf("arr3: %T\n", arr3)

	arr2[2] = "C"
	fmt.Println(arr2)

	fmt.Println(len(arr1))
}
