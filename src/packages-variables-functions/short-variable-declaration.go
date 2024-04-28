package main

/**
在函数中, 短赋值语句 := 可以隐式确定类型的var声明中使用
函数外的每个语句都必须以关键字开始， 因此:=结构不能在函数外使用
*/

import "fmt"

func main() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no"
	fmt.Println(i, j, k, c, python, java)
}
