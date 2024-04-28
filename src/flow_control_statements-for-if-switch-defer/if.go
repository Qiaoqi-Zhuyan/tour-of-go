package main

/**
Go的if语句与for循环类似, 表达式外无需小括号()
而大括号{}则是必须的
*/

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
	//用于将输入的数据转换为字符串。
	//这个函数接受任意数量的参数，
	//并将它们转换为相应的字符串表示，然后将这些字符串连接起来返回。
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
}
