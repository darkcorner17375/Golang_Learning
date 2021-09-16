package main

import "fmt"

//進階: 宣告結構體
type rect struct {
	width, height int
}

//長方形面積計算 - 指標型
func (r *rect) area() int {
	return r.width * r.height
}

//周常計算方法
func (r rect) perim() int {
	return 2 * (r.width + r.height)
}

func main() {
	//變數調用
	r := rect{width: 10, height: 4}
	fmt.Println("area: ", r.area())
	fmt.Println("perim: ", r.perim())

	//指標調用
	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim: ", rp.perim())
}
