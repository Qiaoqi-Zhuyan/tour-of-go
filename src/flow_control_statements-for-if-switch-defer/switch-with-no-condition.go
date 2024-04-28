package main

/**
无条件的 switch 同 switch true 一样。

这种形式能将一长串 if-then-else 写得更加清晰。
*/
import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	if true {

	}

	switch {
	case t.Hour() < 12:
		fmt.Println("good morning")
	case t.Hour() < 17:
		fmt.Println("good afternoon")
	default:
		fmt.Println("good evening")
	}
}
