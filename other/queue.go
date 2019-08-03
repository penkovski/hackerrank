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

type queue struct {
	head *node
	tail *node
}

func (q *queue) empty() bool {
	return q.head == nil
}

func (q *queue) peek() int {
	return q.head.value
}

func (q *queue) add(value int) {
	node := &node{value: value}
	if q.tail == nil {
		q.tail = node
		q.head = node
		return
	}

	q.tail.next = node
	q.tail = node
}

func (q *queue) remove() int {
	if q.head == nil {
		return -1
	}

	v := q.head.value
	q.head = q.head.next
	if q.head == nil {
		q.tail = nil
	}

	return v
}

func (q *queue) String() string {
	var buf bytes.Buffer
	for e := q.head; e != nil; e = e.next {
		s := strconv.Itoa(e.value)
		buf.Write([]byte(s))
		if e.next != nil {
			buf.Write([]byte(" - "))
		}
	}
	return buf.String()
}

func main() {
	q := &queue{}
	fmt.Println(q)

	q.add(5)
	fmt.Println(q)

	q.add(3)
	fmt.Println(q)

	q.add(4)
	fmt.Println(q)

	q.remove()
	fmt.Println(q)

	q.remove()
	fmt.Println(q)

	q.add(10)
	fmt.Println(q)
}
