package main

import (
	"fmt"
	"math"
)

/*
可以为指针类型的接收者声明方法。
这意味着对于某类型 T，接收者的类型可以用 *T 的文法。
此外，T 本身不能是指针，比如不能是 *int。

例如，这里为 *Vertex 定义了 Scale 方法。
指针接收者的方法可以修改接收者指向的值（如这里的 Scale 所示）
由于方法经常需要修改它的接收者，指针接收者比值接收者更常用。
若使用值接收者，那么 Scale 方法会对原始 Vertex 值的副本进行操作。
（对于函数的其它参数也是如此。）
Scale 方法必须用指针接收者来更改 main 函数中声明的 Vertex 的值。
*/

type Vertex_pointers struct {
	X, Y float64
}

func (v Vertex_pointers) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex_pointers) Scale(f float64) {
	v.X = v.Y * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex_pointers{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}
