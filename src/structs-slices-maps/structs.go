package main

/**
一个结构体 struct 就是一组字段 filed
*/

import "fmt"

type Vertex struct {
	x int
	y int
}

func main() {
	fmt.Println(Vertex{1, 2})
}
