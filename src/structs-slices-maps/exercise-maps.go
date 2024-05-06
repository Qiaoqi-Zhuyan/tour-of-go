package main

import (
	"fmt"
	"strings"
)

/**
实现 WordCount。
它应当返回一个映射，其中包含字符串 s 中每个“单词”的个数。
你会发现 strings.Fields 很有用。
*/

func WordCount(s string) map[string]int {
	wc_maps := make(map[string]int)
	words := strings.Fields(s)
	for _, word := range words {
		_, ok := wc_maps[word]
		if ok {
			wc_maps[word]++
		} else {
			wc_maps[word] = 1
		}
	}
	return wc_maps
}

func main() {
	s := "package main import ( golang.org/x/tour/wc ) func WordCount(s string) map[string]int { return map[string]int{ x : 1} } func main() { wc.Test(WordCount) }"
	wc := WordCount(s)
	for i, word := range wc {
		fmt.Println("words: ", i, " value: ", word)
	}
}
