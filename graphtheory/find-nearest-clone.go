package main

import (
	"fmt"
)

var nodes []*node

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	nodes = make([]*node, n)
	for i := 0; i < n; i++ {
		nodes[i] = &node{}
	}

	for i := 0; i < m; i++ {
		var edge1, edge2 int
		fmt.Scan(&edge1, &edge2)

		node1 := nodes[edge1-1]
		node2 := nodes[edge2-1]

		node1.edges = append(node1.edges, node2)
		node2.edges = append(node2.edges, node1)
	}

	for i := 0; i < n; i++ {
		var color int
		fmt.Scan(&color)
		nodes[i].color = color
	}

	var analyzecolor int
	fmt.Scan(&analyzecolor)

	fmt.Println(solve(analyzecolor))
}

func solve(color int) int {
	fmt.Println("analyze color = ", color)

	shortest := 0
	queue := &queue{}
	for i := 0; i < len(nodes); i++ {
		if nodes[i].tried || nodes[i].color != color {
			continue
		}
		nodes[i].tried = true

		path := 0
		found := false
		queue.add(nodes[i])
		for node := queue.remove(); node != nil; {
			fmt.Println("path =", path)
			fmt.Printf("[queue] node = %# v\n", node)
			if path > 0 && node.color == color {
				found = true
				break
			}

			path++
			for _, n := range node.edges {
				queue.add(n)
			}
		}

		fmt.Println("found =", found)
		fmt.Println("shortest =", shortest)
		fmt.Println("path = ", path)

		if found {
			if shortest == 0 || shortest > path {
				shortest = path
			}
		}
	}

	return shortest
}

type node struct {
	color int
	edges []*node
	tried bool
}

type queuenode struct {
	data *node
	next *queuenode
}

type queue struct {
	head *queuenode
	tail *queuenode
}

func (q *queue) empty() bool {
	return q.head == nil
}

func (q *queue) peek() *node {
	return q.head.data
}

func (q *queue) add(n *node) {
	qnode := &queuenode{data: n}
	if q.tail == nil {
		q.tail = qnode
		q.head = qnode
		return
	}

	q.tail.next = qnode
	q.tail = qnode
}

func (q *queue) remove() *node {
	if q.head == nil {
		return nil
	}

	v := q.head.data
	q.head = q.head.next
	if q.head == nil {
		q.tail = nil
	}

	return v
}
