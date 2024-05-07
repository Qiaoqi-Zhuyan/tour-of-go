package main

import (
	"fmt"
	"math"
)

/**
反之也一样：
接受一个值作为参数的函数必须接受一个指定类型的值：
```go
var v Vertex
fmt.Println(AbsFunc(v))  // OK
fmt.Println(AbsFunc(&v)) // 编译错误！
```
而以值为接收者的方法被调用时，接收者既能为值又能为指针：

```go
var v Vertex
fmt.Println(v.Abs()) // OK
p := &v
fmt.Println(p.Abs()) // OK
```
这种情况下，方法调用 p.Abs() 会被解释为 (*p).Abs()。
*/

type Vertex_ind struct {
	X, Y float64
}

func (v Vertex_ind) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Vertex_ind) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex_ind{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))

	p := &Vertex_ind{3, 4}
	fmt.Println(p.Abs())
	fmt.Println(AbsFunc(*p))
}
