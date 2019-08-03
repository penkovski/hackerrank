package main

import (
	"fmt"
)

var graph []*node

func main() {
	var q int
	fmt.Scan(&q)

	for i := 0; i < q; i++ {
		var n, m int
		fmt.Scan(&n, &m)

		// create all nodes
		graph = make([]*node, n)
		for i := 0; i < n; i++ {
			graph[i] = &node{weight: 0}
		}

		// fill up their relationships
		for j := 0; j < m; j++ {
			var u, v int
			fmt.Scan(&u, &v)

			node1, node2 := graph[u-1], graph[v-1]
			node1.siblings = append(node1.siblings, v-1)
			node2.siblings = append(node2.siblings, u-1)
		}

		// get the start node
		var start int
		fmt.Scanln(&start)

		bfsWeights(start)

		for j := 0; j < len(graph); j++ {
			if j == start-1 {
				continue
			}
			if graph[j].visited {
				fmt.Print(graph[j].weight)
			} else {
				fmt.Print("-1")
			}
			fmt.Print(" ")
		}

		fmt.Println()

		//fmt.Printf("graph = %# v\n", pretty.Formatter(graph))
	}
}

type node struct {
	visited  bool
	weight   int
	siblings []int
}

func bfsWeights(start int) {
	var queue []*node
	queue = append(queue, graph[start-1])
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		node.visited = true
		if node == graph[start-1] {
			node.weight = 0
		}

		for _, child := range node.siblings {
			if !graph[child].visited {
				graph[child].visited = true
				graph[child].weight += node.weight + 6
				queue = append(queue, graph[child])
			}
		}
	}
}
