package main

import "fmt"

/**
从之前的练习中复制 Sqrt 函数，修改它使其返回 error 值。

Sqrt 接受到一个负数时，应当返回一个非 nil 的错误值。复数同样也不被支持。
创建一个新的类型
```go
type ErrNegativeSqrt float64
```
并为其实现
```go
func (e ErrNegativeSqrt) Error() string
```
方法使其拥有 error 值，
通过 ErrNegativeSqrt(-2).Error()
调用该方法应返回 "cannot Sqrt negative number: -2"。

修改 Sqrt 函数，使其接受一个负数时，返回 ErrNegativeSqrt 值。

*/

func Sqrt(x float64) (float64, error) {
	return 0, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
