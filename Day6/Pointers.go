package main

import "fmt"

//接收一個數值
func byValue(ival int) {
	ival = 0
}

//些收一個指向int的指標
func byPointer(iptr *int) {
	*iptr = 0
}

//指標調用
func main() {
	//宣告一個數值
	i := 10
	fmt.Println("initial:", i)
	//傳遞給函數byValue
	byValue(i)
	fmt.Println("byvalue:", i)
	//地址傳遞給函數byValue
	byPointer(&i)
	fmt.Println("byptr:", i)
	//打印出i的地址
	fmt.Println("pointer:", &i)

}
