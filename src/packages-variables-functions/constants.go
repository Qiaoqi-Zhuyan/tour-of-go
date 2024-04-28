package main

/**
常量的声明与变量类似
只不过使用const 关键字
常量可以是字符、字符串、布尔值或数值
常量不能用 := 语法声明
*/

import "fmt"

const Pi = 3.1415926

func main() {
	const WORLD = "world"
	fmt.Println("hello", WORLD)
	fmt.Println("Happy", Pi, "Day")

	const TRUTH = true
	fmt.Println("GO rules", TRUTH)
}
