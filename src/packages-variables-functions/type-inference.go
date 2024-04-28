package main

/**
在声明一个变量而不指定其类型时 即使用不带类型的 := 语法 var = 表达式语法
变量的类型会通过右值推断出来
当声明的右值确定了类型时, 新变量的类型与其相同
```go
var i int
j := i
```
不过当右边包含未指明类型的数值常量的时候
新变量的类型就可能是int, float64或者complex128
取决于常量的精度
```go
	i := 42
	f := 3.142
	g := 0.867 + 0.5i
```
*/

import "fmt"

func main() {
	v := 42
	fmt.Printf("v is type of %T", v)
}
