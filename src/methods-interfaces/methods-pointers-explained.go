package main

/**
现在我们要把 Abs 和 Scale 方法重写为函数。
同样，先试着移除掉第 16 的 *，你能看出程序行为改变的原因吗？
要怎样做才能让该示例顺利通过编译？
*/
import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func Abs2(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	Scale(&v, 10)
	fmt.Println(Abs2(v))
}
