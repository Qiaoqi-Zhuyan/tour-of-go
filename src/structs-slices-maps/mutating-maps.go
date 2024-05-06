package main

/**
在映射 m 中插入或修改元素
```go
m[key] = elem
```
获取元素
```
elem = m[key]
```
删除元素
```go
delete(m, key)
```
通过双赋值检测某个键是否存在
```go
elem, ok = m[key]
```
若 key 在 m 中，ok 为 true ；否则，ok 为 false。若 key 不在映射中，则 elem 是该映射元素类型的零值。

注：若 elem 或 ok 还未声明，你可以使用短变量声明：

elem, ok := m[key]


*/

import "fmt"

func main() {
	m := make(map[string]int)

	m["answer"] = 42
	fmt.Println("value: ", m["answer"])

	m["answer"] = 48
	fmt.Println("value: ", m["answer"])

	delete(m, "answer")
	fmt.Println("value: ", m["answer"])

	v, ok := m["answer"]
	fmt.Println("value: ", v, "exist? : ", ok)
}
