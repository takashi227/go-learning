package main

import "fmt"

func integers() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	ints := integers()
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())

	ints2 := integers()
	fmt.Println(ints2())
	fmt.Println(ints2())
	fmt.Println(ints2())
}
