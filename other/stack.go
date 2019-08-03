package main

import (
	"bytes"
	"fmt"
	"strconv"
)

type node struct {
	value int
	next  *node
}

type stack struct {
	head *node
}

func (s *stack) empty() bool {
	return s.head == nil
}

func (s *stack) peek() int {
	return s.head.value
}

func (s *stack) push(value int) {
	node := &node{value: value}
	if s.head == nil {
		s.head = node
		return
	}

	head := s.head
	s.head = node
	s.head.next = head
}

func (s *stack) pop() int {
	if s.head == nil {
		return -1
	}

	v := s.head.value
	s.head = s.head.next
	return v
}

func (s *stack) String() string {
	var buf bytes.Buffer
	for e := s.head; e != nil; e = e.next {
		buf.Write([]byte(strconv.Itoa(e.value)))
		if e.next != nil {
			buf.Write([]byte(" - "))
		}
	}
	return buf.String()
}

func main() {
	stack := &stack{}
	fmt.Println(stack.String())

	stack.push(5)
	fmt.Println(stack.String())

	stack.push(7)
	fmt.Println(stack.String())

	stack.push(9)
	fmt.Println(stack.String())

	stack.pop()
	fmt.Println(stack.String())

	stack.push(15)
	fmt.Println(stack.String())
}
