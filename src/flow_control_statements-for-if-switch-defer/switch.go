package main

/**
GO的switch语句类似于C, C++, Java
不过GO只会运行选定的case, 而非之后所有的case，而非之后所有的case
在效果上, go的做法相当于这些语言中为每个case后面自动添加了所需的break
在GO中，除非以fallthrough语句结束, 否则分支会自动终止
GO的switch的case无需为常量, 且取值不限于整数
*/

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("GO sys environment")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("macOS")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s.\n\n", os)
	}
}
