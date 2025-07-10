package syntax

import "fmt"

var (
	u8 uint8 = 255
	u16 uint16 = 65535
	u32 uint32 = 4294967295
	u64 uint64 = 18446744073709551615
	i8 int8 = 127
	i16 int16 = 32767
	i32 int32 = 2147483647
)

func Int() {
	fmt.Println(u8, u16, u32, u64, i8, i16, i32)
	fmt.Printf("%T %v\n", u8, u8)

	// 演算
	x := 1 + 1
	fmt.Println(x)
	
	y := 0
	y++
	fmt.Println(y)

	y--
	fmt.Println(y)

	fmt.Println("1 + 1 =", 1+1)
	fmt.Println("10 - 1 =", 10-1)
	fmt.Println("10 / 2 =", 10/2)
	fmt.Println("10 * 2 =", 10*2)
	fmt.Println("10 % 2 =", 10%2)

	fmt.Println(10 / 3)
	fmt.Println(10 % 3)
	fmt.Println(10.0 / 3.0)

	// シフト演算
	fmt.Println(1 << 0) // 0001 0001
	fmt.Println(1 << 1) // 0001 0010
	fmt.Println(1 << 2) // 0001 0100
	fmt.Println(1 << 3) // 0001 1000
}