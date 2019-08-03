package main

import "fmt"

var freqs map[int]int

func main() {
	var q int64
	fmt.Scanln(&q)

	freqs = make(map[int]int, q/2)

	var i int64
	for i = 0; i < q; i++ {
		var o, n int
		fmt.Scan(&o)
		fmt.Scan(&n)

		//fmt.Println(q, o, n)
		//return

		switch o {
		case 1:
			freqs[n]++
		case 2:
			if freqs[n] > 0 {
				freqs[n]--
			}
		case 3:
			found := 0
			checked := make(map[int]bool)
			for k, v := range freqs {
				if checked[k] {
					continue
				}
				checked[k] = true
				if v == n {
					found++
				}
			}
			if found > 0 {
				fmt.Println(1)
			} else {
				fmt.Println(0)
			}
		}
	}
}
