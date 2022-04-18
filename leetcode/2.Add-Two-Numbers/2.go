package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	nodes := &ListNode{Val: 0}
	pointer := nodes
	p1 := l1
	p2 := l2
	carry := 0

	for p1 != nil || p2 != nil {
		sum := carry

		if p1 != nil {
			sum += p1.Val
			p1 = p1.Next
		}

		if p2 != nil {
			sum += p2.Val
			p2 = p2.Next
		}

		carry = sum / 10
		quotient := sum % 10

		pointer.Next = &ListNode{Val: quotient}
		pointer = pointer.Next
	}

	if carry != 0 {
		pointer.Next = &ListNode{Val: carry}
	}

	return nodes.Next
}
