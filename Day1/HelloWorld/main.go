package main

import "fmt"

func main() {

	//第一個程式碼
	fmt.Println("Hello DackCorNer world!!")

	//fmt.Print 不會換行
	fmt.Print("A")
	fmt.Print("B")
	fmt.Print("C")

	fmt.Print("\n")

	//fmt.Println 會換行
	fmt.Println("A")
	fmt.Println("B")
	fmt.Println("C")

	//多個輸出
	fmt.Print("A", "B", "C")
	fmt.Print("\n")
	fmt.Println("A", "B", "C")

	//fmt.Printf 格式化輸出
	var a int = 10
	var b int = 5
	var c int = 2
	// fmt.Println("a=", a, "b=", b, "c=", c)
	fmt.Printf("a=%v,b=%v,c=%v", a, b, c)
}
