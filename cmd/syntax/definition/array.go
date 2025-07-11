package definition

import "fmt"

func Array() {
	var a [2]int	
	a[0] = 100
	a[1] = 200
	fmt.Println(a)

	var b [2]int = [2]int{100, 200}
	// MEMO: 配列は固定長なので、appendできない
	// b = append(b, 300)
	fmt.Println(b)

	var c = []int{100, 200}
	c = append(c, 300)
	fmt.Println(c)
}