package main

import "fmt"

/**
go 只有一种循环结构， for 循环
基本的for循环由三部分组成, 用分号隔开
	初始化语句
	条件表达式
	后置语句
初始化语句通常为一句短变量声明, 该变量声明仅在for语句的作用域中可见
和c, java语言不同的是
GO语言的for语句后面的三个构成部分外没有小括号, 大括号是必须有的
*/

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	sum_2 := 1
	for sum_2 < 1000 {
		sum_2 += sum_2
	}
	fmt.Println(sum_2)
}
