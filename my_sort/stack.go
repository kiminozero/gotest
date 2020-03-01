package mysort

import (
	"fmt"
)

//stack by []int
type stack struct {
	S    []int
	topS int
}

type Stack interface {
	StackEmpty() bool
	Push(x int) error
	Pop() (int, error)
}

// double stack  by one []int. if topL + 1 == topR than stack full. 
type dStack struct {
	S    []int
	topL int
	topR int
}

type DStack interface {
	StackEmpty() bool
	LPush(x int) error
	RPush(x int) error
	LPop() (int, error)
	RPop() (int, error)
}

//stack by linklist
type (
	Node struct {
		val  int
		next *Node
	}

	stackL struct {
		S    *Node
		topS int
	}
)

func NewStack(size int) *stack {
	s := stack{make([]int, size), 0}
	fmt.Println("NewStack ", s)
	return &s
}
func (s *stack) StackEmpty() bool {
	if s.topS == 0 {
		return true
	}
	return false
}

func (s *stack) Push(x int) error {
	if s.topS >= len(s.S) {
		return fmt.Errorf("overflow")
	}
	s.S[s.topS] = x
	s.topS++
	return nil
}

func (s *stack) Pop() (int, error) {
	if s.StackEmpty() {
		return 0, fmt.Errorf("underflow")
	} else {
		s.topS--
		return s.S[s.topS], nil
	}
}

func NewStackL() *stackL {
	s := stackL{&Node{0, nil}, 0}
	fmt.Println("NewStackL ", s)
	return &s
}
func (s *stackL) StackEmpty() bool {
	if s.topS == 0 {
		return true
	}
	return false
}

func (s *stackL) Push(x int) error {
	cur := &Node{x, s.S}
	s.S = cur
	s.topS++
	return nil
}

func (s *stackL) Pop() (int, error) {
	if s.StackEmpty() {
		return 0, fmt.Errorf("underflow")
	} else {
		s.topS--
		cur := s.S
		s.S = s.S.next
		return cur.val, nil
	}
}
