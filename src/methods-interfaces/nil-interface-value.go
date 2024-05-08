package main

import "fmt"

/**
nil 接口值既不保存值也不保存具体类型。

为 nil 接口调用方法会产生运行时错误
因为接口的元组内并未包含能够指明该调用哪个 具体 方法的类型。
*/

type I_nil_inter interface {
	M_nil_inter()
}

func main() {
	var i I_nil_inter
	describe_nil_interfaces(i)
	i.M_nil_inter()
}

func describe_nil_interfaces(i I_nil_inter) {
	fmt.Printf("(%v, %T)\n", i, i)

}
