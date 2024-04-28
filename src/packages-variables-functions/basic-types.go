package main

/**
GO 的基本类型
bool
string
int int8 int16 int32 int64 uint uint8 uint16 uint32 uint64 uintptr
byte(uint8)
rune(int32 - unicode)
float32 float64
complex64 complex126

和导入语句一样，变量声明也可以进行分组成一个代码块
int、uint 和 uintptr 类型在 32-位系统上通常为 32-位宽，
在 64-位系统上则为 64-位宽。
当你需要一个整数值时应使用 int 类型，
除非你有特殊的理由使用固定大小或无符号的整数类型。
*/

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("type: %T value: %v\n", ToBe, ToBe)
	fmt.Printf("type: %T value: %v\n", MaxInt, MaxInt)
	fmt.Printf("type: %T value: %v\n", z, z)
}
