package main

import "fmt"

func main() {

	//範例1
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	//Break 印出0.1.2.3.4
	i2 := 0
	for {
		fmt.Println(i2)
		i2++
		if i2 == 5 {
			break
		}
	}

	//continue 印出1.3.5.7.9
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

	//陣列迴圈
	s := []int{1, 2, 3}
	for k, v := range s {
		fmt.Println(k, v)
	}

	//高雄郵遞區號
	PostCode := map[string]int{
		"新興區": 801,
		"前金區": 802,
		"苓雅區": 803,
		"鹽埕區": 804,
		"鼓山區": 805,
		"旗津區": 806,
		"前鎮區": 807,
	}
	//如果想回傳區號，可以用底線忽略地區
	for _, v := range PostCode {
		fmt.Println(v)
	}
}
