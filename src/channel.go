package main

import (
	"fmt"
	"time"
)

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

	// チャネルのクローズ
	ch4 := make(chan int, 2)
	close(ch4)

	// 閉じたチャネルにはデータを送信できない
	// ch4 <- 1
	// 閉じたチャネルからデータを受信できる？
	// fmt.Println(<-ch4)

	// 閉じたチャネルからデータを受信するとfalseになる
	// チャネルが空でクローズ状態だとfalseになる
	i, ok := <-ch4
	fmt.Println(i, ok)

	ch5 := make(chan int, 2)
	go reciever("1.goroutin", ch5)
	go reciever("2.goroutin", ch5)
	go reciever("3.goroutin", ch5)

	i5 := 0
	for i5 < 100 {
		ch5 <- i5
		i5++
	}
	close(ch5)
	time.Sleep(3 * time.Second)

	// for
	ch6 := make(chan int, 10)
	ch6 <- 1
	ch6 <- 2
	ch6 <- 3
	close(ch6) // for文でrangeを使う場合はクローズする必要がある
	for i6 := range ch6 {
		fmt.Println(i6)
	}

	// select
	ch7 := make(chan int, 2)
	ch8 := make(chan string, 2)

	ch8 <- "A"
	ch7 <- 1
	ch8 <- "B"
	ch7 <- 2

	// v1 := <-ch7
	// v2 := <-ch8
	// fmt.Println(v1)
	// fmt.Println(v2)

	select {
	case v1 := <-ch7:
		fmt.Println(v1 + 1000)
	case v2 := <-ch8:
		fmt.Println(v2 + "!!!")
	default:
		fmt.Println("どちらでもない")
	}

	ch9 := make(chan int)
	ch10 := make(chan int)
	ch11 := make(chan int)

	go func() {
		for {
			i9 := <-ch9
			ch10 <- i9 * 2
		}
	}()

	go func() {
		for {
			i10 := <-ch10
			ch11 <- i10 - 1
		}
	}()

	n := 0
L:
	for {
		select {
		case ch9 <- n:
			n++
		case i11 := <-ch11:
			fmt.Println("recieved", i11)
		default:
			if n > 100 {
				break L
			}
		}
		// if n > 100 {
		// 	break
		// }
	}
}

func reciever(name string, ch chan int) {
	for {
		i6, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(name, i6)
	}
	fmt.Println(name + " is done.")
}
