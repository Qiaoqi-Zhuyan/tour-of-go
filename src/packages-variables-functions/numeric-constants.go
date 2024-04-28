package main

/**
数值常量是高精度的值, 一个未指定类型的常量由上下文决定其类型
*/

import "fmt"

const (
	BIG   = 1 << 100 // 1 左移100位是1后面跟着100个0的二进制
	SMALL = BIG >> 99
)

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(needInt(SMALL))
	fmt.Println(needFloat(SMALL))
	fmt.Println(needFloat(BIG))
}
