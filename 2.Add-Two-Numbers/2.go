package main

import (
	"fmt"
	"math"
)

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

func (l *ListNode) iterateNode() {
	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}
}

func makeListNode(nums []int) *ListNode {
	nodes := &ListNode{Val: -1}
	pointer := nodes
	for _, num := range nums {
		if nodes.Val == -1 {
			nodes.Val = num
		} else {
			nodes.Next = &ListNode{Val: num}
			nodes = nodes.Next
		}
	}

	return pointer
}

func getLenNodes(l *ListNode) int {
	length := 0
	for l != nil {
		length++
		l = l.Next
	}
	return length
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	numOfListOne := getLenNodes(l1)
	numOfListTwo := getLenNodes(l2)

	max := int(math.Max(float64(numOfListOne), float64(numOfListTwo)))

	nodes := &ListNode{Val: -1}
	result := nodes
	p1 := l1
	p2 := l2
	carry := 0
	for i := 0; i < max; i++ {
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

		if nodes.Val == -1 {
			nodes.Val = quotient
		} else {
			nodes.Next = &ListNode{Val: quotient, Next: nil}
			nodes = nodes.Next
		}
	}

	if carry != 0 {
		nodes.Next = &ListNode{Val: carry}
	}

	return result
}

func main() {
	list1 := []int{9, 9, 9, 9, 9, 9, 9}
	list2 := []int{9, 9, 9, 9}
	l1 := makeListNode(list1)
	l2 := makeListNode(list2)

	node := addTwoNumbers(l1, l2)
	node.iterateNode()
}
