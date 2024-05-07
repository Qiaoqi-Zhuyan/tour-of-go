package main

import "fmt"

/**
比较前两个程序，你大概会注意到带指针参数的函数必须接受一个指针：

```go
var v Vertex
ScaleFunc(v, 5)  // 编译错误！
ScaleFunc(&v, 5) // OK
```
而接收者为指针的的方法被调用时
接收者既能是值又能是指针：
```go
var v Vertex
v.Scale(5)  // OK
p := &v
p.Scale(10) // OK
```
对于语句 v.Scale(5) 来说，即便 v 是一个值而非指针
，带指针接收者的方法也能被直接调用。
也就是说，由于 Scale 方法有一个指针接收者，
为方便起见，Go 会将语句 v.Scale(5) 解释为 (&v).Scale(5)。
*/

type Vertex_indirection struct {
	X, Y float64
}

func (v *Vertex) Scale_indirection(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale_indirection(2)
	ScaleFunc(&v, 10)

	p := &Vertex{4, 3}
	p.Scale_indirection(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)

}
