package main

import (
	"fmt"
	"math"
)

/**
接口类型 的定义为一组方法签名。
接口类型的变量可以持有任何实现了这些方法的值

 示例代码的第 22 行存在一个错误。
由于 Abs 方法只为 *Vertex （指针类型）定义，
因此 Vertex（值类型）并未实现 Abser。
*/

type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat_interfaces(-math.Sqrt2) // a MyFloat 实现了 Abser
	v := Vertex_interfaces{3, 4}         // a *Vertex 实现了 Abser

	a = f
	fmt.Println(a.Abs())

	a = &v
	fmt.Println(a.Abs())

}

type MyFloat_interfaces float64

func (f MyFloat_interfaces) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex_interfaces struct {
	X, Y float64
}

func (v *Vertex_interfaces) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y + v.Y)
}
