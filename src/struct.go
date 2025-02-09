package main

import "fmt"

// User型の構造体を定義
type User struct {
	// フィールドを定義
	Name string
	Age  int
	// X, Y int
}

func UpdateUser(user User) {
	user.Name = "A"
	user.Age = 1000
}

func UpdateUser2(user *User) {
	user.Name = "B"
	user.Age = 2000
}

func main() {
	// User型の変数を宣言
	var user1 User
	fmt.Println(user1)

	// フィールドに値を代入
	user1.Name = "User1"
	user1.Age = 30
	fmt.Println(user1)

	user2 := User{}
	fmt.Println(user2)

	// フィールドに値を代入
	user2.Name = "User2"
	user2.Age = 20
	fmt.Println(user2)

	// フィールド名を指定して初期化
	user3 := User{Name: "User3", Age: 10}
	fmt.Println(user3)

	// フィールド名を指定せずに初期化
	// 定義してあるフィールドの順番で初期値を代入する必要がある
	user4 := User{"User4", 40}
	fmt.Println(user4)

	user5 := User{Name: "User5"}
	fmt.Println(user5)

	// new関数を使って構造体のポインタを取得
	user6 := new(User)
	fmt.Println(user6)

	// アドレス演算子を使って構造体のポインタを取得
	user7 := &User{}
	fmt.Println(user7)

	UpdateUser(user1)
	UpdateUser2(user7)
	fmt.Println(user1)
	fmt.Println(user7)
}
