package main

import (
	"fmt"

	"github.com/cami-crane/lesson-go.git/cmd/syntax/definition"
)

// 最初に実行される関数
func init() {
	fmt.Println("init")
}

func main() {
	definition.HelloWorld()
	definition.Import()
	definition.Variabledefinition()
	definition.Const()
	definition.Int()
	definition.String()
	definition.Bool()
	definition.Cast()
	definition.Array()
	definition.Slice()
	definition.MakeAndCap()
	definition.Map()
	definition.Bite()
}
