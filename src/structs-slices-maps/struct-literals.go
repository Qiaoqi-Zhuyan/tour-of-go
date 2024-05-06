package main

/**
使用Name: 语法可以仅列出部分字段 字段名的顺序无关
特殊前缀&返回一个只想结构体的指针
*/

import "fmt"

type Vertex3 struct {
	X, Y int
}

var (
	v1 = Vertex3{1, 2}
	v2 = Vertex3{X: 1} // Y: 被隐式赋值于0
	v3 = Vertex3{}
	p  = &Vertex3{1, 2} // 创建一个*Vertex 类型的结构体指针
)

func main() {
	fmt.Println(v1, p, v2, v3)
}
