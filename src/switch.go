package main

import "fmt"

func anything(a interface{}) {
	// fmt.Println(a)
	switch v := a.(type) {
	case int:
		fmt.Println(v + 10000)
	case string:
		fmt.Println(v + "!?")
	}
}

func main() {
	var n int = 1
	switch n {
	case 1, 2:
		fmt.Println("1 or 2")
	case 3, 4:
		fmt.Println("3 or 4")
	default:
		fmt.Println("unknown")
	}

	switch n := 2; n {
	case 1, 2:
		fmt.Println("1 or 2")
	case 3, 4:
		fmt.Println("3 or 4")
	default:
		fmt.Println("unknown")
	}

	var n2 int = 6
	switch {
	case n2 > 0 && n2 < 4:
		fmt.Println("0 < n < 4")
	case n2 > 3 && n2 < 7:
		fmt.Println("3 < n < 7")
	}

	anything(1)
	anything("aaa")

	var x interface{} = "test"
	i, isInt := x.(int)
	fmt.Println(i+2, isInt)

	f, isFloat64 := x.(float64)
	fmt.Println(f, isFloat64)

	if x == nil {
		fmt.Println("None")
	} else if i, isInt := x.(int); isInt {
		fmt.Println(i, "x is int")
	} else if s, isString := x.(string); isString {
		fmt.Println(s, "x is string")
	} else {
		fmt.Println("I don't know")
	}

	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("I don't know")
	}

	switch v := x.(type) {
	case bool:
		fmt.Println(v, "bool")
	case int:
		fmt.Println(v, "int")
	case string:
		fmt.Println(v, "string")
	default:
		fmt.Println("I don't know")
	}
}
