package main

import (
	"fmt"
	"sort"
)

type Node struct {
	id       int
	weight   int
	visited  bool
	children []*Node
}

func main() {
	var q int // number of queries
	fmt.Scanln(&q)

	for i := 0; i < q; i++ {
		var nodes = make(map[int]*Node)

		var n int // number of nodes
		var m int // number of edges
		fmt.Scan(&n, &m)

		for j := 0; j < m; j++ {
			var n1, n2 int
			fmt.Scan(&n1, &n2)

			// create both nodes if they don't exist
			if _, ok := nodes[n1]; !ok {
				nodes[n1] = &Node{id: n1, weight: 0, visited: false, children: make([]*Node, n)}
			}

			if _, ok := nodes[n2]; !ok {
				nodes[n2] = &Node{id: n2, weight: 0, visited: false, children: make([]*Node, n)}
			}

			node1 := nodes[n1]
			node2 := nodes[n2]

			node1.children = append(node1.children, node2)
			node2.children = append(node2.children, node1)
		}

		var startNode int
		fmt.Scan(&startNode)

		bfs(startNode, nodes)

		// fill in the gaps for nodes not mentioned in the queries
		for k := 1; k <= n; k++ {
			if _, ok := nodes[k]; !ok {
				nodes[k] = &Node{id: k, weight: 0, visited: false, children: make([]*Node, n)}
			}
		}

		// sort nodes by id (key)
		var keys []int
		for k := range nodes {
			keys = append(keys, k)
		}
		sort.Ints(keys)

		// print all distances
		for _, k := range keys {
			if k == startNode {
				continue
			}

			n := nodes[k]
			if n.weight <= 0 {
				fmt.Print("-1")
			} else {
				fmt.Print(n.weight)
			}
			fmt.Print(" ")
		}

		fmt.Println() // new line for the next query to start clean
	}
}

func bfs(startNode int, nodes map[int]*Node) {
	//	fmt.Printf("[bfs] startNode = %d, nodes = %v\n", startNode, nodes)

	var queue []*Node
	start := nodes[startNode]
	queue = append(queue, start)

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node != nil {
			node.visited = true

			for _, child := range node.children {
				if child == nil || child.visited {
					continue
				}

				child.visited = true
				child.weight += node.weight + 6
				queue = append(queue, child)
			}
		}
	}
}

