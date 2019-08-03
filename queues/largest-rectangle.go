package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanln(&n)

	var heights []int
	for i := 0; i < n; i++ {
		var h int
		fmt.Scan(&h)
		heights = append(heights, h)
	}

	fmt.Println(largestRectangle(heights))
}

func largestRectangle(heights []int) int {

}

type queue struct {
	head *node
	tail *node
}

type node struct {
	value int
	next  *node
}

func (q *queue) add(value int) {
	node := &node{value: value}
	if q.head == nil {
		q.head = node
		q.tail = node
		return
	}

	q.tail.next = node
	q.tail = node
}

func (q *queue) remove() {

}
