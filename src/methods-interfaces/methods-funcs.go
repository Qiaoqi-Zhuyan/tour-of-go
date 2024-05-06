package main

import (
	"fmt"
	"math"
)

type Vertex_method_funcs struct {
	X, Y float64
}

func Abs(v Vertex_method_funcs) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex_method_funcs{3, 4}
	fmt.Println(Abs(v))
}
