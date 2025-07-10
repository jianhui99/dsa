// https://leetcode.com/problems/remove-linked-list-elements/description/

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
	head = [1,2,6,3,4,5,6], val = 6

output:
	result: [1,2,3,4,5]
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	node := &ListNode{Next: head}
	curr := node

	for curr.Next != nil {
		if curr.Next.Val == val {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}

	return node.Next
}

func PrintListNode(head *ListNode) {
	cur := head
	for cur != nil {
		fmt.Printf("[%d] -> ", cur.Val)
		cur = cur.Next
	}
	fmt.Println("nil")
}

func main() {
	ln := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 1,
						Next: &ListNode{
							Val: 1,
							Next: &ListNode{
								Val: 1,
							},
						},
					},
				},
			},
		},
	}

	r := removeElements(ln, 1)
	PrintListNode(r)
}
