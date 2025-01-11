package main

import "fmt"

var i3 int = 300

func outer() {
	var s2 string = "outer"
	fmt.Println(s2)
}

func main() {
	var i int = 100
	fmt.Println(i)

	var i2 int
	var s string
	fmt.Println(i2, s)

	i2 = 200
	s = "GO"
	fmt.Println(i2, s)

	fmt.Println(i3)

	outer()
}
