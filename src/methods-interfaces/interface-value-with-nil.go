package main

import "fmt"

/**
即便接口内的具体值为 nil，方法仍然会被 nil 接收者调用。
在一些语言中，这会触发一个空指针异常
但在 Go 中通常会写一些方法来优雅地处理它（如本例中的 M 方法）。
保存了 nil 具体值的接口其自身并不为 nil。
*/

type I_nil interface {
	M_nil()
}

type T_nil struct {
	S string
}

func (t *T_nil) M_nil() {
	if t == nil {
		fmt.Println("nil")
		return
	}
	fmt.Println(t.S)
}

func main() {
	var i I_nil
	var t *T_nil

	i = t
	describe_nil(i)
	i.M_nil()

	i = &T_nil{"hello"}
	describe_nil(i)
	i.M_nil()
}

func describe_nil(i I_nil) {
	fmt.Printf("(%v, %T)\n", i, i)
}
