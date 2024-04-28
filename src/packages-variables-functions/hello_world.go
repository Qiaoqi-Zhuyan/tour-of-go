package main

import (
	"fmt"
	"math/rand"
)

import "math"

func add(x int, y int) int {
	return x + y
}

func add_(x, y int) int {
	return x + y
}

// 带名字的返回值, 能够反映其含义
/**
没有参数的return语句会直接返回已命名的返回值，也就是裸值
*/
func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println("hello world", rand.Intn(10))
	fmt.Printf("hello world %g\n", math.Sqrt(7))
	fmt.Println(math.Pi)
	fmt.Println(add(42, 13))
	fmt.Println(add_(42, 13))
	a, b := swap("world", "hello")
	fmt.Println(a, b)
	y, x := split(17)
	fmt.Println(x, y)
	fmt.Println(split(17))
}
