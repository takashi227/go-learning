package main

import "fmt"

type MyInt int

func (mi MyInt) Print() {
	fmt.Println(mi)
}

func main() {
	var mi MyInt
	fmt.Println(mi)
	fmt.Printf("%T\n", mi)

	// MyInt型とint型では演算できない
	// i := 1
	// fmt.Println(mi + i)

	mi.Print()
}
