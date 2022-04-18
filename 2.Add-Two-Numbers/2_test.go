package main_test

import (
	"fmt"
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

func Test_Problem(t *testing.T) {
	tc := []testCase{
		{
			Input{[]int{}, []int{}},
			Answer{[]int{}},
		},
	}

	for _, c := range tc {
		expected := c.Answer
		acutal := main.addTwoNumbers(c.Input.ListOne, c.Input.ListTwo)
		assert.
	}
}
