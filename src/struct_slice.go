package main

import "fmt"

type User struct {
	Name string
	Age  int
}

type Users []*User

// type Users struct {
// 	Users []*User
// }

func main() {
	user1 := User{Name: "User1", Age: 10}
	user2 := User{Name: "User2", Age: 20}
	user3 := User{Name: "User3", Age: 30}
	user4 := User{Name: "User4", Age: 40}
	user5 := User{Name: "User5", Age: 50}

	users := Users{}
	users = append(users, &user1)
	users = append(users, &user2)
	users = append(users, &user3, &user4, &user5)
	// users = append(users, &user4)
	// users = append(users, &user5)
	// fmt.Println(users)

	for _, user := range users {
		fmt.Println(*user)
	}

	users2 := make([]*User, 0)
	users2 = append(users2, &user1, &user2, &user3)

	for _, user := range users2 {
		fmt.Println(*user)
	}
}
