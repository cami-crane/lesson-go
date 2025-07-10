package main

import (
	"fmt"

	"github.com/cami-crane/lesson-go.git/cmd/syntax"
)

// 最初に実行される関数
func init() {
	fmt.Println("init")
}

func main() {
	syntax.HelloWorld()
	syntax.Import()
	syntax.VariableDefinition()
	syntax.Const()
	syntax.Int()
	syntax.String()
}
