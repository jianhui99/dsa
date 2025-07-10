// https://leetcode.com/problems/remove-duplicates-from-sorted-list/description/

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

func deleteDuplicates(head *ListNode) *ListNode {
	current := head

	for current != nil && current.Next != nil {
		if current.Val == current.Next.Val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}

	return head
}

func printList(head *ListNode) {
	cur := head
	for cur != nil {
		fmt.Printf("[%d] -> ", cur.Val)
		cur = cur.Next
	}
	fmt.Println("nil")
}

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 3,
					},
				},
			},
		},
	}
	m := deleteDuplicates(head)
	printList(m)

}
