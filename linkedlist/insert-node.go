package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	fmt.Scan(&n)

	var head *node
	var curr *node
	for i := 0; i < n; i++ {
		var data int
		fmt.Scanln(&data)
		if head == nil {
			head = &node{data: data}
			curr = head
			continue
		}

		curr.next = &node{data: data}
		curr = curr.next
	}

	var newvalue int
	fmt.Scanln(&newvalue)
	var newposition int
	fmt.Scanln(&newposition)

	insertAtPosition(head, newvalue, newposition)

	fmt.Println(head.String())
}

type node struct {
	data int
	next *node
}

func (n *node) String() string {
	if n == nil {
		return ""
	}
	return strconv.Itoa(n.data) + " " + n.next.String()
}

func insertAtPosition(list *node, data int, position int) *node {
	newnode := &node{data: data}
	if position == 0 {
		newnode.next = list
		return newnode
	}

	currpos := 0
	for n := list; n != nil; n = n.next {
		if currpos == position-1 {
			newnode.next = n.next
			n.next = newnode
			break
		}
		currpos++
	}

	return list
}
