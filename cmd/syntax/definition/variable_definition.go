package definition

import "fmt"

func separatedefinition() {
	var i int = 1
	var f64 float64 = 1.2
	var s string = "test"
	var t bool = true
	var f bool = false

	// 複数定義も可
	var t2, f2 bool = true, false

	fmt.Println(i, f64, s, t, f, t2, f2)
}

// var は関数外でも宣言できる	
var (
	i int = 1
	f64 float64 = 1.2
	s string = "test"
	t bool = true
	f bool = false
)

func collectivelyDefined() {
	fmt.Println(i, f64, s, t, f)
}

// var は関数外では宣言できない
func shortdefinition() {
	i := 1
	f64 := 1.2
	s := "test"
	t := true
	f := false
	
	fmt.Println(i, f64, s, t, f)
}

func typedefinition() {
	var xi int = 1
	// デフォルトでは float64 になる
	var xf32 float32 = 1.2
	var xs string = "test"
	var xt, xf bool = true, false
	
	fmt.Println(xi, xf32, xs, xt, xf)
	// 型を確認する
	// 勝手には改行されないので \n を入れる
	fmt.Printf("%T\n", xf32)
	fmt.Printf("%T\n", xf32)
}
func Variabledefinition() {
	separatedefinition()
	collectivelyDefined()
	shortdefinition()
	typedefinition()
}