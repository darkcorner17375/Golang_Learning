package main

import (
	"fmt"
	"math"
)

const message string = "2021轉職之路"

func main() {

	fmt.Println(message)
	//message = "不能更改"

	//常量運算
	const a = 100
	const b = 200 / a
	fmt.Println(a, b)
	fmt.Println(int64(a / b))

	//math中的常量
	fmt.Println(math.Pi)
	fmt.Println(math.Sin(math.Pi / 6))
	fmt.Println(math.Cos(math.Pi / 3))
	fmt.Println(math.Tan(math.Pi / 4))

}
