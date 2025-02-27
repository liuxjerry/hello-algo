// File: deque_test.go
// Created Time: 2022-11-29
// Author: Reanon (793584285@qq.com)

package chapter_stack_and_queue

import (
	"container/list"
	"fmt"
	"testing"

	. "github.com/krahets/hello-algo/pkg"
)

func TestDeque(t *testing.T) {
	/* 初始化双向队列 */
	// 在 Go 中，将 list 作为双向队列使用
	deque := list.New()

	/* 元素入队 */
	deque.PushBack(2)
	deque.PushBack(5)
	deque.PushBack(4)
	deque.PushFront(3)
	deque.PushFront(1)
	fmt.Print("双向队列 deque = ")
	PrintList(deque)

	/* 访问元素 */
	front := deque.Front()
	fmt.Println("队首元素 front =", front.Value)
	rear := deque.Back()
	fmt.Println("队尾元素 rear =", rear.Value)

	/* 元素出队 */
	deque.Remove(front)
	fmt.Print("队首出队元素 front = ", front.Value, "，队首出队后 deque = ")
	PrintList(deque)
	deque.Remove(rear)
	fmt.Print("队尾出队元素 rear = ", rear.Value, "，队尾出队后 deque = ")
	PrintList(deque)

	/* 获取双向队列的长度 */
	size := deque.Len()
	fmt.Println("双向队列长度 size =", size)

	/* 判断双向队列是否为空 */
	isEmpty := deque.Len() == 0
	fmt.Println("双向队列是否为空 =", isEmpty)
}

func TestLinkedListDeque(t *testing.T) {
	// 初始化队列
	deque := NewLinkedListDeque()

	// 元素入队
	deque.OfferLast(2)
	deque.OfferLast(5)
	deque.OfferLast(4)
	deque.OfferFirst(3)
	deque.OfferFirst(1)
	fmt.Print("队列 deque = ")
	PrintList(deque.toList())

	// 访问队首元素
	front := deque.PeekFirst()
	fmt.Println("队首元素 front =", front)
	rear := deque.PeekLast()
	fmt.Println("队尾元素 rear =", rear)

	// 元素出队
	pollFirst := deque.PollFirst()
	fmt.Print("队首出队元素 pollFirst = ", pollFirst, "，队首出队后 deque = ")
	PrintList(deque.toList())
	pollLast := deque.PollLast()
	fmt.Print("队尾出队元素 pollLast = ", pollLast, "，队尾出队后 deque = ")
	PrintList(deque.toList())

	// 获取队的长度
	size := deque.Size()
	fmt.Println("队的长度 size =", size)

	// 判断是否为空
	isEmpty := deque.IsEmpty()
	fmt.Println("队是否为空 =", isEmpty)
}

// BenchmarkArrayQueue 67.92 ns/op in Mac M1 Pro
func BenchmarkLinkedListDeque(b *testing.B) {
	stack := NewLinkedListDeque()
	// use b.N for looping
	for i := 0; i < b.N; i++ {
		stack.OfferLast(777)
	}
	for i := 0; i < b.N; i++ {
		stack.PollFirst()
	}
}
