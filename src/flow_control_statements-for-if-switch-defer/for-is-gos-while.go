package main

/**
去掉分号, C的while在Go中叫做for
*/

import "fmt"

func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
