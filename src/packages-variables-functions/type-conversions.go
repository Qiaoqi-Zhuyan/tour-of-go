package main

/**
表达式 T(v) 将值 v 转换为类型 T
一些数值类型的转化
'''go
	var i int = 42
	var f float64 = float64(i)
	var u uint = uint(f)
'''
或者:
'''go
	i := 42
	f := float64(i)
	u := uint(f)
'''

与C不同的是, GO在不同类型的项之间赋值需要显示转化,
*/

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}
