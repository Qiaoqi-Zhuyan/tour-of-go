package main

import (
	"fmt"
	"math/rand"
)

type Tree struct {
	Letf  *Tree
	Value int
	Right *Tree
}

func insert(t *Tree, v int) *Tree {
	if t == nil {
		return &Tree{
			nil,
			v,
			nil,
		}
	}
	if v < t.Value {
		t.Letf = insert(t.Letf, v)
	} else {
		t.Right = insert(t.Right, v)
	}
	return t
}

func New(k int) *Tree {
	var t *Tree
	for _, v := range rand.Perm(10) {
		t = insert(t, (1+v)*k)
	}
	return t
}

func Walk(t *Tree, ch chan int) {
	if t != nil {
		Walk(t.Letf, ch)
		ch <- t.Value
		Walk(t.Right, ch)
	}
}

func Same(t1, t2 *Tree) bool {
	c1 := make(chan int, 10)
	c2 := make(chan int, 10)
	Walk(t1, c1)
	go Walk(t2, c2)
	for i := 0; i < len(c1); i++ {
		if <-c1 != <-c2 {
			return false
		}
	}
	return true
}

func main() {
	ch := make(chan int, 10)
	tree := New(1)
	Walk(tree, ch)
	//for i := 0; i < cap(ch); i++ {
	//	fmt.Println(<-ch)
	//}

	t1 := New(2)
	t2 := New(3)
	flag := Same(t1, t2)
	fmt.Println(flag)
}
