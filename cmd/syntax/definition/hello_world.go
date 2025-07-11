package definition

import "fmt"

// 文字連結
func buzz() {
	fmt.Println("fizz", "buzz")
}

// 外部インポートはパスカルケース
func HelloWorld() {
	buzz()
	fmt.Println("Hello, world!")
}
