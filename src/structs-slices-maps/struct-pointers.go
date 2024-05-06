package main

/**
结构体字段可通过结构体指针来访问

如果有一个指向结构体的指针p, 那么可以通过(*p).X 来访问
其字段X, 但是也允许我们使用隐式解引用，直接写p.X 就可以

*/

import "fmt"

type Vertex2 struct {
	X int
	Y int
}

func main() {
	v := Vertex2{1, 2}
	p := &v
	p.X = 1e9
	(*p).Y = 2e6
	fmt.Println(v)
}
