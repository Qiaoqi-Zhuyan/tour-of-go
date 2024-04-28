package main

/**
推迟调用的函数调用会被压入一个栈中，当外层函数返回时, 被推迟的调用
会按照先进后出的顺序调用
*/

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("10")
}
