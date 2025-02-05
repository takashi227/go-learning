package main

import "fmt"

// channel
// 複数のゴルーチン間でデータの受け渡しを行うためのデータ構造

func main() {
	// チャネルの宣言（双方向）
	var ch1 chan int

	// 受信専用のチャネル
	// var ch2 <-chan int
	// 送信専用のチャネル
	// var ch3 chan<- int

	// チャネルの初期化
	ch1 = make(chan int)
	ch2 := make(chan int)

	fmt.Println(cap(ch1))
	fmt.Println(cap(ch2))

	// バッファサイズを指定したチャネルの初期化
	ch3 := make(chan int, 5)
	fmt.Println(cap(ch3))

	// チャネルにデータを送信
	ch3 <- 1
	fmt.Println(len(ch3))
	ch3 <- 2
	ch3 <- 3
	fmt.Println("len", len(ch3))

	// チャネルからデータを受信
	i := <-ch3
	fmt.Println(i)
	fmt.Println("len", len(ch3))
	i2 := <-ch3
	fmt.Println(i2)
	fmt.Println("len", len(ch3))
	fmt.Println(<-ch3)
	fmt.Println("len", len(ch3))

	// バッファサイズを超えた場合はデッドロックが発生する
	ch3 <- 1
	fmt.Println("len", len(ch3))
	fmt.Println(<-ch3)
	fmt.Println("len", len(ch3))
	ch3 <- 2
	ch3 <- 3
	ch3 <- 4
	ch3 <- 5
	ch3 <- 6 // デッドロック
}
