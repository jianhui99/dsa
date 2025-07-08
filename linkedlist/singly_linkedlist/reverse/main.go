package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type SinglyLinkedList struct {
	Head *ListNode
	Size int
}

// 反转链表的核心算法
func (ll *SinglyLinkedList) Reverse() {
	var prev *ListNode // 前一个节点，初始为nil
	current := ll.Head // 当前节点，从头开始

	for current != nil {
		next := current.Next // 第1步：保存下一个节点
		current.Next = prev  // 第2步：反转当前节点的指向
		prev = current       // 第3步：移动prev指针
		current = next       // 第4步：移动current指针
	}

	ll.Head = prev // 最后更新头指针
}

// 带详细注释的反转函数，用于演示
func reverseWithDebug(head *ListNode) *ListNode {
	fmt.Println("=== 开始反转链表 ===")

	var prev *ListNode = nil // 前驱节点
	current := head          // 当前节点
	step := 1

	for current != nil {
		fmt.Printf("第%d步开始:\n", step)
		fmt.Printf("  prev: %v\n", getNodeValue(prev))
		fmt.Printf("  current: %v\n", getNodeValue(current))
		fmt.Printf("  current.Next: %v\n", getNodeValue(current.Next))

		// 核心的4个步骤
		next := current.Next // 1. 保存下一个节点
		fmt.Printf("  next: %v (保存)\n", getNodeValue(next))

		current.Next = prev // 2. 反转当前节点的指向
		fmt.Printf("  反转指向: current.Next = prev\n")

		prev = current // 3. 移动prev指针
		current = next // 4. 移动current指针
		fmt.Printf("  移动指针: prev=%v, current=%v\n", getNodeValue(prev), getNodeValue(current))

		fmt.Println("  ---")
		step++
	}

	fmt.Printf("反转完成，新的头节点: %v\n", getNodeValue(prev))
	return prev
}

// 辅助函数：获取节点值用于打印
func getNodeValue(node *ListNode) string {
	if node == nil {
		return "nil"
	}
	return fmt.Sprintf("[%d]", node.Val)
}

// 打印链表
func printList(head *ListNode) {
	if head == nil {
		fmt.Println("空链表")
		return
	}

	current := head
	for current != nil {
		fmt.Printf("[%d]", current.Val)
		if current.Next != nil {
			fmt.Print(" -> ")
		}
		current = current.Next
	}
	fmt.Println()
}

// 创建测试链表
func createTestList() *ListNode {
	// 创建 [1] -> [2] -> [3] -> nil
	node1 := &ListNode{Val: 1, Next: nil}
	node2 := &ListNode{Val: 2, Next: nil}
	node3 := &ListNode{Val: 3, Next: nil}

	node1.Next = node2
	node2.Next = node3

	return node1
}

func main() {
	// 创建原始链表 [1] -> [2] -> [3]
	head := createTestList()
	fmt.Println("原始链表:")
	printList(head)

	fmt.Println()

	// 执行反转并显示详细过程
	newHead := reverseWithDebug(head)

	fmt.Println("\n反转后的链表:")
	printList(newHead)

	fmt.Println("\n=== 图解反转过程 ===")
	fmt.Println("原始: [1] -> [2] -> [3] -> nil")
	fmt.Println()
	fmt.Println("步骤1: prev=nil, current=[1], next=[2]")
	fmt.Println("       反转后: nil <- [1]    [2] -> [3] -> nil")
	fmt.Println("       移动指针: prev=[1], current=[2]")
	fmt.Println()
	fmt.Println("步骤2: prev=[1], current=[2], next=[3]")
	fmt.Println("       反转后: nil <- [1] <- [2]    [3] -> nil")
	fmt.Println("       移动指针: prev=[2], current=[3]")
	fmt.Println()
	fmt.Println("步骤3: prev=[2], current=[3], next=nil")
	fmt.Println("       反转后: nil <- [1] <- [2] <- [3]    current=nil")
	fmt.Println("       移动指针: prev=[3], current=nil")
	fmt.Println()
	fmt.Println("结束: current=nil，循环结束")
	fmt.Println("      新的头节点: prev=[3]")
	fmt.Println("      最终结果: [3] -> [2] -> [1] -> nil")
}
