package main

import (
	"fmt"
	"math"
)

/**
Go 没有类。不过你可以为类型定义方法。
方法就是一类带特殊的 接收者 参数的函数
方法接收者在它自己的参数列表内，位于 func 关键字和方法名之间。

在此例中，Abs 方法拥有一个名字为 v，类型为 Vertex 的接收者。
*/

type Vertex_method struct {
	X, Y float64
}

func (v Vertex_method) Abs(a int) float64 {
	return math.Sqrt(v.X*v.X+v.Y*v.Y) + float64(a)
}

func main() {
	v := Vertex_method{3, 4}
	fmt.Println(v.Abs(3))
}
