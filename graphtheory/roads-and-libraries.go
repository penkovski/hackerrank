package main

import (
	"fmt"
)

func main() {
	var q int
	fmt.Scan(&q)

	for i := 0; i < q; i++ {
		var n, m, croad, clib int
		fmt.Scan(&n, &m, &clib, &croad)

		cities := make([]*city, n)
		for i := 0; i < n; i++ {
			cities[i] = &city{}
		}

		// gather roads
		for j := 0; j < m; j++ {
			var road1, road2 int
			fmt.Scan(&road1, &road2)

			index1, index2 := road1-1, road2-1

			city1, city2 := cities[index1], cities[index2]
			city1.id, city2.id = index1, index2

			city1.roads = append(city1.roads, index2)
			city2.roads = append(city2.roads, index1)
		}

		//fmt.Printf("cities = %# v\n", pretty.Formatter(cities))
		//fmt.Println("road cost =", croad, " library cost =", clib)

		fmt.Println(solve(cities, croad, clib))
	}
}

func solve(cities []*city, croad int, clib int) int {
	if croad >= clib {
		return clib * len(cities)
	}

	cost := 0
	for i := 0; i < len(cities); i++ {
		//fmt.Printf("[solve] i = %d, len(cities) = %d\n", i, len(cities))
		//fmt.Printf("[solve] cities[%d].visited = %v\n", i, cities[i].visited)

		if !cities[i].visited {
			cost += clib

			visitedCities := cities[i].visit(cities)
			//fmt.Printf("[solve] visitedCities = %# v\n", pretty.Formatter(visitedCities))
			if visitedCities > 1 {
				cost += visitedCities*croad - croad
			}
		}
	}
	return cost
}

type city struct {
	id      int
	roads   []int
	visited bool
}

func (c *city) visit(cities []*city) int {
	//fmt.Printf("[visit] city = %# v\n", pretty.Formatter(c))

	cities[c.id].visited = true
	if len(c.roads) == 0 {
		return 1
	}

	visitiedCities := 1
	for i := 0; i < len(c.roads); i++ {
		if cities[c.roads[i]].visited {
			continue
		}
		visitiedCities += cities[c.roads[i]].visit(cities)
	}
	return visitiedCities
}
