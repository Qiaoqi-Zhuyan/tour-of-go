package main

import (
	"fmt"
	"math"
)

/**
可以为非结构体类型声明方法
看到了一个带 Abs 方法的数值类型 MyFloat
你只能为在同一个包中定义的接收者类型声明方法
而不能为其它别的包中定义的类型 （包括 int 之类的内置类型）声明方法。

就是接收者的类型定义和方法声明必须在同一包内。
*/

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
