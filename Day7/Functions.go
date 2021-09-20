package main

import "fmt"

func main() {
	sayMessage("Hello GO!")

	//指標呼叫
	greeting := "Hello"
	name := "Edward"
	sayGreeting(&greeting, &name)
	fmt.Println(name)

	//return 回傳
	s := sum(1, 2, 3, 4, 5)
	fmt.Println("The sum is", s)
}

func sayMessage(msg string) {
	fmt.Println(msg)
}

func sayGreeting(greeting, name *string) {
	fmt.Println(*greeting, *name)
	*name = "Ted"
	fmt.Println(*name)
}
func sum(values ...int) (result int) {
	fmt.Println(values)
	for _, v := range values {
		result += v
	}
	return
}
