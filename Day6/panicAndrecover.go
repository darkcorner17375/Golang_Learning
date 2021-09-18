package main

import "fmt"

func main() {
	fmt.Println("Hello Golang!")

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
		fmt.Println("有疫苗在別怕")
	}()
	panic("發現多名確診者!")
	// fmt.Println("是否會執行")
}
