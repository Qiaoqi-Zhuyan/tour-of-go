package main

/**
switch的case语句从上到下顺序执行, 直到匹配成功时停止
```go
switch i{
	case 0:
	case f():
}
```
在 i==0, f不会被调用

*/

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Saturday is ")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("today")
	case today + 1:
		fmt.Println("tomorrow")
	case today + 2:
		fmt.Println("the day after tomorrow")
	default:
		fmt.Println("many days later")
	}
}
