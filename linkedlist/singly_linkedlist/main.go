package main

import (
	"fmt"
	"strings"
)

// 单链表节点结构
type ListNode struct {
	Val  int
	Next *ListNode
}

// 单链表结构
type SinglyLinkedList struct {
	Head *ListNode
	Size int
}

// 创建新的单链表
func NewSinglyLinkedList() *SinglyLinkedList {
	return &SinglyLinkedList{
		Head: nil,
		Size: 0,
	}
}

// 在链表末尾添加节点
func (ll *SinglyLinkedList) Append(val int) {
	newNode := &ListNode{Val: val, Next: nil}

	if ll.Head == nil {
		ll.Head = newNode
	} else {
		current := ll.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newNode
	}
	ll.Size++
}

// 在链表开头添加节点
func (ll *SinglyLinkedList) Prepend(val int) {
	newNode := &ListNode{Val: val, Next: ll.Head}
	ll.Head = newNode
	ll.Size++
}

// 在指定位置插入节点
func (ll *SinglyLinkedList) Insert(index int, val int) error {
	if index < 0 || index > ll.Size {
		return fmt.Errorf("index out of range")
	}

	if index == 0 {
		ll.Prepend(val)
		return nil
	}

	newNode := &ListNode{Val: val, Next: nil}
	current := ll.Head
	for i := 0; i < index-1; i++ {
		current = current.Next
	}

	newNode.Next = current.Next
	current.Next = newNode
	ll.Size++
	return nil
}

// 删除第一个值为val的节点
func (ll *SinglyLinkedList) Delete(val int) bool {
	if ll.Head == nil {
		return false
	}

	if ll.Head.Val == val {
		ll.Head = ll.Head.Next
		ll.Size--
		return true
	}

	current := ll.Head
	for current.Next != nil {
		if current.Next.Val == val {
			current.Next = current.Next.Next
			ll.Size--
			return true
		}
		current = current.Next
	}

	return false
}

// 查找值为val的节点，返回索引
func (ll *SinglyLinkedList) Find(val int) int {
	current := ll.Head
	index := 0

	for current != nil {
		if current.Val == val {
			return index
		}
		current = current.Next
		index++
	}

	return -1
}

// 获取指定位置的节点值
func (ll *SinglyLinkedList) Get(index int) (int, error) {
	if index < 0 || index >= ll.Size {
		return 0, fmt.Errorf("index out of range")
	}

	current := ll.Head
	for i := 0; i < index; i++ {
		current = current.Next
	}

	return current.Val, nil
}

// 反转链表
func (ll *SinglyLinkedList) Reverse() {
	var prev *ListNode
	current := ll.Head

	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	ll.Head = prev
}

// 显示链表
func (ll *SinglyLinkedList) Display() string {
	if ll.Head == nil {
		return "空链表"
	}

	var result []string
	current := ll.Head

	for current != nil {
		result = append(result, fmt.Sprintf("[%d]", current.Val))
		current = current.Next
	}

	return strings.Join(result, " -> ")
}

// 获取链表长度
func (ll *SinglyLinkedList) Length() int {
	return ll.Size
}

// 检查链表是否为空
func (ll *SinglyLinkedList) IsEmpty() bool {
	return ll.Size == 0
}

// 清空链表
func (ll *SinglyLinkedList) Clear() {
	ll.Head = nil
	ll.Size = 0
}

// 创建示例链表 [1] -> [2] -> [3]
func createExampleList() *SinglyLinkedList {
	ll := NewSinglyLinkedList()
	ll.Append(1)
	ll.Append(2)
	ll.Append(3)
	return ll
}

func main() {
	// 创建您提到的链表
	linkedList := createExampleList()
	fmt.Println("原始链表:", linkedList.Display())

	fmt.Println("\n=== 链表操作演示 ===")

	// 在开头插入
	linkedList.Prepend(0)
	fmt.Println("在开头插入0:", linkedList.Display())

	// 在末尾添加
	linkedList.Append(4)
	fmt.Println("在末尾添加4:", linkedList.Display())

	// 在指定位置插入
	err := linkedList.Insert(2, 15)
	if err != nil {
		fmt.Println("插入失败:", err)
	} else {
		fmt.Println("在位置2插入15:", linkedList.Display())
	}

	// 查找元素
	index := linkedList.Find(3)
	fmt.Printf("查找元素3的位置: %d\n", index)

	// 获取指定位置的元素
	value, err := linkedList.Get(3)
	if err != nil {
		fmt.Println("获取失败:", err)
	} else {
		fmt.Printf("位置3的元素值: %d\n", value)
	}

	// 删除元素
	deleted := linkedList.Delete(15)
	if deleted {
		fmt.Println("删除15后:", linkedList.Display())
	} else {
		fmt.Println("删除15失败")
	}

	// 反转链表
	fmt.Println("反转前:", linkedList.Display())
	linkedList.Reverse()
	fmt.Println("反转后:", linkedList.Display())

	fmt.Printf("\n链表长度: %d\n", linkedList.Length())
	fmt.Printf("链表是否为空: %t\n", linkedList.IsEmpty())

	// 清空链表
	linkedList.Clear()
	fmt.Println("清空后:", linkedList.Display())
	fmt.Printf("链表长度: %d\n", linkedList.Length())
}
