package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func (u User) SayName() {
	fmt.Println(u.Name)
}

func (u User) SetName(name string) {
	u.Name = name
}

func (u *User) SetName2(name string) {
	u.Name = name
}

func main() {
	user1 := User{Name: "User1", Age: 10}
	user1.SayName()

	user1.SetName("A")
	user1.SayName()

	user1.SetName2("B")
	user1.SayName()

	user2 := &User{Name: "User2", Age: 20}
	user2.SayName()

	user2.SetName2("C")
	user2.SayName()
}
