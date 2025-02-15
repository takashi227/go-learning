package main

import (
	"fmt"
	"time"
)

func main() {
	// 時間の生成
	// 今の時間
	t := time.Now()
	fmt.Println(t)

	// 指定した時間を生成
	t2 := time.Date(2020, 2, 10, 10, 10, 10, 0, time.Local)
	fmt.Println(t2)
	fmt.Println(t2.Year())
	fmt.Println(t2.YearDay())
	fmt.Println(t2.Month())
	fmt.Println(t2.Weekday())
	fmt.Println(t2.Day())
	fmt.Println(t2.Hour())
	fmt.Println(t2.Minute())
	fmt.Println(t2.Second())
	fmt.Println(t2.Nanosecond())
	fmt.Println(t2.Zone())

	// 時間の感覚を表現する
	// Time.Duration型を返す
	fmt.Println(time.Hour)
	fmt.Printf("%T\n", time.Hour)
	fmt.Println(time.Minute)
	fmt.Println(time.Second)
	fmt.Println(time.Millisecond)
	fmt.Println(time.Microsecond)
	fmt.Println(time.Nanosecond)

	// time.Duration型を文字列から生成する
	d, _ := time.ParseDuration("1h30m")
	fmt.Println(d)

	t3 := time.Now()
	t3 = t3.Add(2*time.Minute + 15*time.Second)
	fmt.Println(t3)

	// 時間の差分を取得
	t4 := time.Date(2025, 2, 10, 0, 0, 0, 0, time.Local)
	t5 := time.Now()
	fmt.Println(t5)

	// t4 - t5
	fmt.Println(t4.Sub(t5))
	// t5 - t4
	fmt.Println(t5.Sub(t4))

	// 時刻を比較する
	fmt.Println(t5.Before(t4))
	fmt.Println(t5.After(t4))
	fmt.Println(t4.Before(t5))
	fmt.Println(t4.After(t5))

	// 指定時間のスリープ
	// 5秒間停止
	time.Sleep(5 * time.Second)
	fmt.Println("5秒経過しました")
}
