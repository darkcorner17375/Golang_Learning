package main

import "fmt"

func main() {

	//宣告變數 a,b
	var a int = 10
	b := 15
	fmt.Println(a, b)

	//宣告多個變數
	var c int = 110
	c, d := 100, 120
	fmt.Println(c, d)

	//宣告字串
	var e = "你好,歡迎!"
	fmt.Println(e)

	//宣告數值
	var f int = 10
	fmt.Println(f)

	//宣告布林
	var g = true
	fmt.Println(g)

	//宣告浮點數
	var h float32 = 3.141592
	fmt.Println(h)

	//指標型
	var x string = "a"
	var pf *string = &x
	//資料的位址
	fmt.Println(pf)
	//資料
	fmt.Println(*pf)
}
