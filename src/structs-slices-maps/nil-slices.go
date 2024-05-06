package main

/**

切片的零值是nil
nil切片长度和容量为0且没有底层数组

*/
import "fmt"

func main() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil")
	}
}
