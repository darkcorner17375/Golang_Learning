package main

import "fmt"

func MyDefer() {
	//延遲調用
	defer fmt.Println("開始")
	fmt.Println("結束")
}

func main() {
	fmt.Println("主程式開始運行")
	MyDefer()
	MyDefer2()
	MyDefer3()
}

func MyDefer2() {
	fmt.Println("處理1")
	defer fmt.Println("處理2")
	fmt.Println("處理3")
}

func MyDefer3() {
	defer fmt.Println("處理4")
	defer fmt.Println("處理5")
	defer fmt.Println("處理6")
}
