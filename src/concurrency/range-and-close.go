package main

import "fmt"

/*
发送者可通过 close 关闭一个信道来表示没有需要发送的值了。
接收者可以通过为接收表达式分配第二个参数来测试信道是否被关闭
若没有值可以接收且信道已被关闭，那么在执行完
```go
v, ok := <-ch
```
此时 ok 会被设置为 false。
循环 for i := range c 会不断从信道接收值，直到它被关闭。

 只应由发送者关闭信道，而不应油接收者关闭。
向一个已经关闭的信道发送数据会引发程序 panic。

 信道与文件不同，通常情况下无需关闭它们。
只有在必须告诉接收者不再有需要发送的值时才有必要关闭
例如终止一个 range 循环。
*/

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	//for i := range c {
	//	fmt.Println(i)
	//}

	for i := 0; i < cap(c); i++ {
		fmt.Println(<-c)
	}
}
