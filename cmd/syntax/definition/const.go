package definition

import "fmt"

const pi = 3.14

// まとめて定義
const (
	Username = "test_user"
	Password = "test_password"
)

// var big int = 9223372036854775807 + 1
const big = 9223372036854775807 + 1

func Const() {
	fmt.Println(pi)
	fmt.Println(big - 1)
}