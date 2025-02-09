package main

import "fmt"

type T struct {
	User
	// Customer
}

type User struct {
	Name string
	Age  int
}

type Customer struct {
	Name string
	Age  int
}

func (u *User) SetName() {
	u.Name = "A"
}

func main() {
	t := T{
		User: User{
			Name: "User1",
			Age:  10,
		},
		// Customer: Customer{
		// 	Name: "Customer1",
		// 	Age:  20,
		// },
	}
	fmt.Println(t)
	fmt.Println(t.User)
	fmt.Println(t.User.Name)
	fmt.Println(t.User.Age)

	fmt.Println(t.Name)

	// t.User.SetName()
	// fmt.Println(t.User.Name)
	t.SetName()
	fmt.Println(t.Name)
}
