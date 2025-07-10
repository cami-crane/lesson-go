package syntax

import (
	"fmt"
	"strings"
)

func String() {
	fmt.Println("Hello World!")
	fmt.Println("Hello" + "World!")
	fmt.Println(string("Hello World"[0]))

	var s string = "Hello World"
	
		// "H" を "X" に置き換える（最初の1つだけ）
		fmt.Println(strings.Replace(s, "H", "X", 1))
		// "World" が含まれているかどうかを確認
		fmt.Println(strings.Contains(s, "World"))
		// "l" の出現回数をカウント
		fmt.Println(strings.Count(s, "l"))
		// "World" の最初の出現位置を探す
		fmt.Println(strings.Index(s, "World"))
		// "l" の最後の出現位置を探す
		fmt.Println(strings.LastIndex(s, "l"))
		// 文字列を大文字に変換
		fmt.Println(strings.ToUpper(s))
		// 文字列を小文字に変換
		fmt.Println(strings.ToLower(s))
}