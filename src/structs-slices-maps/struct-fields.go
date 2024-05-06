package main

/**
结构体字段可通过点号, 来访问
*/

import "fmt"

type Vertex_ struct {
	X int
	Y int
}

func main() {
	v := Vertex_{1, 2}
	v.X = 4
	fmt.Println(v.X)
}
