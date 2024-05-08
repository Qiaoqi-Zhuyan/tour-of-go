package main

import "fmt"

/**
指定了零个方法的接口值被称为 *空接口：*
```go
interface{}
```
空接口可保存任何类型的值。
（因为每个类型都至少实现了零个方法。）
空接口被用来处理未知类型的值。
例如，fmt.Print 可接受类型为 interface{} 的任意数量的参数。
*/

func main() {
	var i interface{}

	i = 42
	describe_empty(i)

	i = "hello"
	describe_empty(i)

	i = &Vertex_empty{3, 4}
	describe_empty(i)
}

func describe_empty(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

type Vertex_empty struct {
	x, y float64
}
