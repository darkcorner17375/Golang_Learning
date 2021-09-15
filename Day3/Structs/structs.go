package main

import "fmt"

//宣告結構體
type user struct {
	name     string
	password string
	age      int
}

func main() {
	//給結構體附值
	user1 := user{name: "Dark", password: "12345678", age: 25}
	fmt.Println(user1.name, user1.password, user1.age)

	//宣告指標 for user1
	user1p := &user1
	fmt.Println(user1p.name, user1p.password, user1p.age)

	//裡用指標給user1改變內容
	user1p.name = "Edward"
	user1p.password = "987654321"
	user1p.age = 18
	fmt.Println(user1.name, user1.password, user1.age)
	fmt.Println(user1p.name, user1p.password, user1p.age)

}
