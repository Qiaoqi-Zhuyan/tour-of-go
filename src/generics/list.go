package main

/**
除了泛型函数之外，Go 还支持泛型类型。
 类型可以使用类型参数进行参数化，
这对于实现通用数据结构非常有用。

此示例展示了能够保存任意类型值的单链表的简单类型声明。

作为练习，请为此链表的实现添加一些功能。


*/

type List[T any] struct {
	next *List[T]
	pre  *List[T]
	val  T
}

func (l *List[T]) init() {
	l.next = nil
	l.pre = nil
}

func (l *List[T]) push_back(v T) {
	node := &List[T]{
		next: nil,
		val:  v,
		pre:  l,
	}
	l.next = node
}

func (l *List[T]) pop() {
	l.pre.next = nil
}
func main() {

}
