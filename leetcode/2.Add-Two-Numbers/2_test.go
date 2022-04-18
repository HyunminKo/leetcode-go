package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	Input
	Answer
}

type Input struct {
	ListOne []int
	ListTwo []int
}

type Answer struct {
	one []int
}

func IntsToListNode(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	nodes := &ListNode{}
	pointer := nodes

	for _, num := range nums {
		pointer.Next = &ListNode{Val: num}
		pointer = pointer.Next
	}

	return nodes.Next
}

func ListNodeToInts(head *ListNode) []int {

	ints := []int{}

	for head != nil {
		ints = append(ints, head.Val)
		head = head.Next
	}

	return ints
}

func Test_Problem(t *testing.T) {
	tc := []testCase{
		{
			Input{[]int{}, []int{}},
			Answer{[]int{}},
		},
		{
			Input{[]int{1}, []int{1}},
			Answer{[]int{2}},
		},
		{
			Input{[]int{1, 2, 3}, []int{1}},
			Answer{[]int{2, 2, 3}},
		},
		{
			Input{[]int{5, 4}, []int{5, 5}},
			Answer{[]int{0, 0, 1}},
		},
		{
			Input{[]int{2, 4, 3}, []int{5, 6, 4}},
			Answer{[]int{7, 0, 8}},
		},
		{
			Input{[]int{9, 9, 9, 9, 9, 9, 9}, []int{9, 9, 9, 9}},
			Answer{[]int{8, 9, 9, 9, 0, 0, 0, 1}},
		},
	}

	for _, c := range tc {
		expected := c.Answer.one

		actual := ListNodeToInts(
			addTwoNumbers(
				IntsToListNode(c.Input.ListOne),
				IntsToListNode(c.Input.ListTwo)))

		assert.Equal(t, expected, actual)
	}
}
