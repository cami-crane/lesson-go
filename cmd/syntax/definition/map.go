package definition

import "fmt"

func Map() {
	m := map[string]int{"apple": 100, "banana": 200}
	fmt.Println(m)
	fmt.Println(m["apple"])
	m["banana"] = 300
	fmt.Println(m)
	m["new"] = 500
	fmt.Println(m)

	fmt.Println(m["nothing"])

	// MEMO: 2つ目の戻り値は指定しなくても良い
	v, ok := m["apple"]
	fmt.Println(v, ok)

	v2, ok2 := m["nothing"]
	fmt.Println(v2, ok2)

	m2 := make(map[string]int)
	m2["pc"] = 5000
	fmt.Println(m2)

	// MEMO: メモリ上に存在せず、panicが発生する
	// var m3 map[string]int
	// m3["pc"] = 5000
	// fmt.Println(m3)

	var s []int
	if s == nil {
		fmt.Println("Nil")
	}
}