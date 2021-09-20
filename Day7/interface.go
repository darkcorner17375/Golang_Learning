package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	//宣告矩形長寬
	r := rect{width: 5, height: 4}
	//宣告圓形半徑
	c := circle{radius: 10}
	//用measure函數計算矩形面積與周長
	measure(r)
	//用measure函數計算圓形面積與周長
	measure(c)
}

//定義幾何圖形interface
type geometry interface {
	area() float64
	perim() float64
}

//定義矩形struct
type rect struct {
	width, height float64
}

//定義圓形struct
type circle struct {
	radius float64
}

//矩形面積
func (r rect) area() float64 {
	return r.width * r.height
}

//矩形周長
func (r rect) perim() float64 {
	return 2 * (r.width + r.height)
}

//圓形面積
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

//圓形周長
func (c circle) perim() float64 {
	return 2 * c.radius * math.Pi
}

//計算函數,參數幾何圖形接口類型
func measure(g geometry) {
	fmt.Println(reflect.TypeOf(g), g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}
