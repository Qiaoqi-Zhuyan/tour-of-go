package main

/**
var 语句用于声明一些列变量, 和函数的参数列表一样，类型在最后
var 可以出现在包或者函数的层级
*/
import "fmt"

var c, python, java bool

func main() {
	var i int
	fmt.Println(i, c, python, java)
}
