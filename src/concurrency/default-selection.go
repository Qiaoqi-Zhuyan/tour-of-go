package main

import (
	"fmt"
	"time"
)

/**
当 select 中的其它分支都没有准备好时，
default 分支就会执行。

为了在尝试发送或者接收时不发生阻塞
可使用 default 分支：

```go
select {
case i := <-c:
    // 使用 i
default:
    // 从 c 中接收会阻塞时执行
}
```
*/

func main() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)

	for {
		select {
		case <-tick:
			fmt.Println("tick")
		case <-boom:
			fmt.Println("boom")
		default:
			fmt.Println(" ")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
