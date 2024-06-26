package main

/**
GO 拥有指针
指针保存了值的内存地址
类型 *T 是指向T类型值的指针, 其零值为nil
```go
var p *int
```
&操作符会生成一个指向其操作数的指针
```go
	i := 42
	p = &i
```
* 操作符表示指针指向的底层值
```go
fmt.Println(*p) // 通过指针 p 读取 i
*p = 21         // 通过指针 p 设置 i
```
这也就是通常所说的「解引用」或「间接引用」。
与 C 不同，Go 没有指针运算。
*/

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         // 指向i
	fmt.Println(*p) // 通过指针读取i的值
	*p = 21         // 通过指针设置i的值
	fmt.Println(i)  // 查看i的值

	p = &j         // 指向j
	*p = *p / 37   // 通过指针对j进行除法运算
	fmt.Println(j) // 查看j
}
