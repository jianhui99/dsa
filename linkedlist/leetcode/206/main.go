// https://leetcode.com/problems/reverse-linked-list/description/

package main

import "fmt"

/**
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
 */

/*
Example
input:
	head : [1] -> [2] -> [2] -> [3] -> [3]

output:
	result: [1] -> [2] -> [3]
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func PrintListNode(head *ListNode) {
	cur := head
	for cur != nil {
		fmt.Printf("[%d] -> ", cur.Val)
		cur = cur.Next
	}
	fmt.Println("nil")
}

func reverse(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head

	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	return prev
}

func reverse2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := reverse2(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead

}

func main() {
	ln := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
			},
		},
	}

	r := reverse2(ln)
	PrintListNode(r)
}
