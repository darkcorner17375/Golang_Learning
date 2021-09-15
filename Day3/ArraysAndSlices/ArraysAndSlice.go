package main

import "fmt"

func main() {
	//宣告陣列
	var a [5]int
	fmt.Println("a陣列:", a)

	//更改陣列值
	a[4] = 100
	fmt.Println("a陣列", a)
	fmt.Println("a[4]", a[4])

	//陣列長度
	fmt.Println("陣列長度", len(a))

	//宣告陣列並給值
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("b陣列為", b)

	//二維陣列
	var c [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			c[i][j] = i + j
		}
	}
	fmt.Println("二維陣列迴圈結果", c)

	//字串陣列
	var students [3]string
	fmt.Printf("Students:%v\n", students)
	students[0] = "Lisa"
	students[1] = "Ahmed"
	students[2] = "Arnold"
	fmt.Printf("Students:%v\n", students)

	//宣告Slice空切片
	s := make([]string, 3)
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("切片內容:", s)
	fmt.Println("s[2]:", s[2])
	fmt.Println("切片長度:", len(s))

	//新增切片內容
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("新增內容後的切片:", s)

	//宣告新的切片並拷貝s內容
	q := make([]string, len(s))
	copy(q, s) // 等同於q := s[:]
	fmt.Println("新切片內容", q)

	//取下切片部分內容
	//取資料2-4
	l := s[2:5]
	fmt.Println("數據234", l)
	//取開始到第四個
	l = s[:5]
	fmt.Println("數據01234", l)
	//取從2開始到結束
	l = s[2:]
	fmt.Println("數據2345", l)
}
