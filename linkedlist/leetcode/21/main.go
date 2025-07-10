// https://leetcode.com/problems/merge-two-sorted-lists/description/

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
	head1 : [1] -> [1] -> [4]
	head2 : [1] -> [2] -> [4] -> [6]

output:
	result: [1] -> [1] -> [1] -> [2] -> [4] -> [4] -> [6]
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	result := &ListNode{}

	current := result

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}

	if list1 != nil {
		current.Next = list1
	} else {
		current.Next = list2
	}

	return result.Next
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
	head1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	head2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 6,
				},
			},
		},
	}

	m := mergeTwoLists(head1, head2)
	printList(m)

}
