package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	m := map[int]User{
		1: User{Name: "User1", Age: 10},
		2: User{Name: "User2", Age: 20},
		3: User{Name: "User3", Age: 30},
		4: User{Name: "User4", Age: 40},
		5: User{Name: "User5", Age: 50},
	}
	fmt.Println(m)

	m2 := map[User]string{
		User{Name: "User1", Age: 10}: "User1",
		User{Name: "User2", Age: 20}: "User2",
		User{Name: "User3", Age: 30}: "User3",
		User{Name: "User4", Age: 40}: "User4",
		User{Name: "User5", Age: 50}: "User5",
	}
	fmt.Println(m2)

	m3 := make(map[int]User)
	fmt.Println(m3)
	m3[1] = User{Name: "User1", Age: 10}
	m3[2] = User{Name: "User2", Age: 20}
	m3[3] = User{Name: "User3", Age: 30}
	fmt.Println(m3)

	for _, user := range m {
		fmt.Println(user)
	}

}
