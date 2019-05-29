package stack

import (
	"fmt"
)

// stack base on  oubly-linked

const (
	defaultCapacity = 16
)

type PoPush interface {
	Pop() (interface{}, error)
	Push(interface{}) (int, error)
}

type Node struct {
	//doubly-linked list
	next, prev *Node

	// The list to which this node belongs.
	stack *Stack

	// The value stored with this node.
	Value interface{}
}

type Stack struct {
	top, bottom    *Node
	capacity, size int
}

func New(capacity int) *Stack {
	s := &Stack{}
	return s.Init(capacity)
}

func (s *Stack) Init(capacity int) *Stack {
	s.top = nil
	s.bottom = nil
	s.size = 0
	if capacity > defaultCapacity {
		s.capacity = capacity
	} else {
		s.capacity = defaultCapacity
	}
	return s
}

func (s *Stack) Empty() bool {
	if s.size == 0 {
		return true
	} else {
		return false
	}
}

func (s *Stack) Full() bool {
	if s.size == s.capacity {
		return true
	} else {
		return false
	}
}

func (s *Stack) Len() int {
	return s.size
}

func (s *Stack) Top() *Node {
	return s.top
}

func (s *Stack) Bottom() *Node {
	return s.bottom
}

func (s *Stack) Push(element interface{}) (int, error) {
	if s.Full() == true {
		return -1, fmt.Errorf("stack is  full,cannot push  any  element")
	}

	node := &Node{next: nil, prev: nil, stack: s, Value: element}
	if s.bottom == nil {
		s.bottom = node
	}
	if s.top == nil {
		s.top = node
	} else {
		s.top.prev = node
		node.next = s.top
	}
	s.size++
	return s.size, nil
}

func (s *Stack) Pop() (interface{}, error) {
	if s.Empty() == true {
		return -1, fmt.Errorf("stack is  empty,cannot pop  any  element")
	}
	var res interface{}
	if s.Len() == 1 {
		res = s.top.Value
		s.top.stack = nil
		s.top = nil
		s.bottom = nil
	} else {
		res = s.top.Value
		s.top.stack = nil
		next := s.top.next
		next.prev = nil
		s.top.next = nil
		s.top = next
		next = nil
	}
	s.size--
	return res, nil

}

func (s *Stack) Exist(element interface{}) bool {
	if s.Empty() == true {
		return false
	}

	for current := s.top; current == nil; current = current.next {
		if current.Value == element {
			return true
		}
	}
	return false
}

func (s *Stack) Size() int {
	return s.size
}

func (s *Stack) Capacity() int {
	return s.capacity
}
