package main

import "fmt"

/**
类型 [n]T 表示一个数组, 拥有n个类型为T的值
表达式
var a [10]int
会将变量a声明为拥有10个整数的数组
数组的长度是其类型的一部分, 因此数组不能改变大小
*/

func main() {

	var a [2]string
	a[0] = "hello"
	a[1] = "world"
	fmt.Println(a[0], a[1])
	fmt.Println(a)
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

}
