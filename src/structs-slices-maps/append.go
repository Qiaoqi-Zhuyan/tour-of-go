package main

import "fmt"

/**
```go
func append(s []T, vs ...T) []T
````
append的第一个参数s是一个元素类型为T的切片,其余类型为T的值将会追加到该切片的末尾
append 的结果是一个包含原切片所有元素加上新添加元素的切片。
当 s 的底层数组太小，不足以容纳所有给定的值时，它就会分配一个更大的数组。
返回的切片会指向这个新分配的数组。
*/

func main() {
	var s []int
	a := []int{1, 2, 3}
	printSlices(s)

	// 在空切片上追加
	s = append(s, 0)
	printSlices(s)

	// 切片会按需增长
	s = append(s, 1)
	printSlices(s)

	// 可以一次性增加多个元素
	s = append(s, 2, 3, 4)
	printSlices(s)

	a = append(a, 1, 2, 3, 4, 5)
	printSlices(a)
}

func printSlices(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
