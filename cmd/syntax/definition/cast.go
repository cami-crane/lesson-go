package definition

import (
	"fmt"
	"strconv"
)

func Cast() {
	var x int = 1
	xx := float64(x)
	fmt.Printf("%T %v %f\n", xx, xx, xx)

	var y float64 = 1.2
	yy := int(y)
	fmt.Printf("%T %v %d\n", yy, yy, yy)

	var z string = "14"
	// MEMO: Goでは下記のような型変換はできない
	// z := int(s)
	i, _ := strconv.Atoi(z)
	fmt.Printf("%T %v %d\n", i, i, i)

	var h string = "Hello"
	fmt.Println(h[0])
	fmt.Println(string(h[0]))
}