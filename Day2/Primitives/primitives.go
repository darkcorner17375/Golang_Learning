package main

import (
	"fmt"
)

func main() {
	//布林值
	var n bool = true
	fmt.Printf("%v , %T\n", n, n)

	n2 := 1 == 1
	n3 := 1 == 2
	fmt.Printf("%v , %T\n", n2, n2)
	fmt.Printf("%v , %T\n", n3, n3)

	//如不給值 預設則False
	var n4 bool
	fmt.Printf("%v , %T\n", n4, n4)

	//精度轉變
	var a int = 10
	var b int8 = 3
	fmt.Println(a + int(b))

	//位元運算子
	a1 := 10              // 1010
	b1 := 3               // 0011
	fmt.Println(a1 & b1)  //0010
	fmt.Println(a1 | b1)  //1011
	fmt.Println(a1 ^ b1)  //1001
	fmt.Println(a1 &^ b1) //0100

	//位元移位運算子
	a2 := 8              // 2^3
	fmt.Println(a2 << 3) // 2^3 * 2^3 = 2^6
	fmt.Println(a2 >> 3) // 2^3 / 2^3 = 2^0

}
