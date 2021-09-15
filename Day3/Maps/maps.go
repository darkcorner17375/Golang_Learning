package main

import "fmt"

func main() {

	//宣告空的key:value 哈希表
	m := make(map[string]int)
	//給值
	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("m:", m)

	//給變數附值
	v1 := m["k1"]
	fmt.Println("v1:", v1)

	//Hash 長度
	fmt.Println("len(m):", len(m))

	//刪除一個Hash
	delete(m, "k2")
	fmt.Println("m:", m)

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
	fmt.Println(PostCode)
	fmt.Println(PostCode["鼓山區"])

	pc := PostCode
	delete(pc, "前鎮區")
	fmt.Println(pc)
	fmt.Println(PostCode)

}
