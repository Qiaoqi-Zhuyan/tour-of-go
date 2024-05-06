package main

/**
切片字面量类似于没有长度的数组字面量

这就是一个数组字面量
```go
[3]bool{true, true, false}
下面这样则会创建一个和上面相同的数组，然后再构建一个引用了他的切片
```
*/

import "fmt"

func main() {
	q := []int{2, 3, 5, 6, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}
