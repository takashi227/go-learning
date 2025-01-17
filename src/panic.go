package main

import "fmt"

func main() {
	defer func() {
		if x := recover(); x != nil {
			fmt.Println(x)
		}
	}()
	panic("パニック発生")
	fmt.Println("START")
}
