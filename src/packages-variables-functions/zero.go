package main

/**
没有明确初始化的变量声明会被赋予对应类型的零值
零值:
	数值类型为 0，
	布尔类型为 false，
	字符串为 ""（空字符串）
*/
import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q", i, f, b, s)
}
