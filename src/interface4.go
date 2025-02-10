package main

import "fmt"

// type Strunger interface {
// 	String() string
// }

type Point struct {
	A int
	B string
}

func (p *Point) String() string {
	return fmt.Sprintf("<<%v, %v>>", p.A, p.B)
}

func main() {
	p := &Point{A: 1, B: "B"}
	fmt.Println(p)
}
