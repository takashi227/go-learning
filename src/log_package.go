package main

import (
	"log"
	"os"
)

func main() {
	// ログの出力先を指定（標準出力）
	log.SetOutput(os.Stdout)

	// ログの出力
	// log.Print("Log\n")
	// log.Println("Log2")
	// log.Printf("Log%d\n", 3)

	// log.Fatal("Log Fatal\n")
	// log.Fatalln("Log2 Fatal")
	// log.Fatalf("Log%d Fatal\n", 3)

	// log.Panic("Log Panic\n")
	// log.Panicln("Log2 Panic")
	// log.Panicf("Log%d Panic\n", 3)

	// 任意のファイルを作成してログを出力
	// f, err := os.Create("test_log.log")
	// if err != nil {
	// 	return
	// }
	// log.SetOutput(f)
	// log.Println("Logをファイルに出力")

	// ログのフォーマットを指定する
	// 標準のログフォーマット
	log.SetFlags(log.LstdFlags)
	log.Println("A")

	// マイクロ秒まで表示
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
	log.Println("B")

	// 時刻とファイルの行番号
	log.SetFlags(log.Ltime | log.Lshortfile)
	log.Println("C")

	// ログのプレフィックスを指定
	log.SetFlags(log.LstdFlags)
	log.SetPrefix("[LOG]")
	log.Println("D")

	// Loggerの生成
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Println("message")

	// 条件分岐、エラーで終了させる
	_, err := os.Open("A.txt")
	if err != nil {
		log.Fatalln(err)
	}

}
