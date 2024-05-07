package main

import (
	"fmt"
	"math"
)

/**
使用指针接收者的原因有二：
首先，方法能够修改其接收者指向的值。
其次，这样可以避免在每次调用方法时复制该值。
若值的类型为大型结构体时，这样会更加高效。

在本例中，Scale 和 Abs 接收者的类型为 *Vertex，即便 Abs 并不需要修改其接收者。
通常来说，所有给定类型的方法都应该有值或指针接收者，
但并不应该二者混用。
*/

type Vertex_receiver struct {
	X, Y float64
}

func (v *Vertex_receiver) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex_receiver) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &Vertex_receiver{3, 4}
	fmt.Printf("before scale: %+v, abs: %v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("after scale: %+v, abs: %v\n", v, v.Abs())
}
