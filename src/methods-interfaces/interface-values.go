package main

import (
	"fmt"
	"math"
)

/**
接口也是值。它们可以像其它值一样传递。
接口值可以用作函数的参数或返回值。

在内部，接口值可以看做包含值和具体类型的元组：
```go
(value, type)
```
接口值保存了一个具体底层类型的具体值。
接口值调用方法时会执行其底层类型的同名方法。
*/

type I_val interface {
	M_val()
}

type T_val struct {
	S string
}

func (t *T_val) M_val() {
	fmt.Println(t.S)
}

type F float64

func (f F) M_val() {
	fmt.Println(f)
}

func main() {
	var i I_val
	i = &T_val{"hello"}
	describe(i)
	i.M_val()

	i = F(math.Pi)
	describe(i)
	i.M_val()
}

func describe(i I_val) {
	fmt.Printf("(%v, %T)\n", i, i)
}
