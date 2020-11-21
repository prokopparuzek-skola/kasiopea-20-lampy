package main

import (
	"fmt"
	"sort"
)

func main() {
	var T int
	fmt.Scan(&T)
	for i := 0; i < T; i++ {
		var N int
		var lamps, spaces []int
		fmt.Scan(&N)
		spaces = make([]int, N-1)
		lamps = make([]int, N)
		for j := 0; j < N; j++ {
			fmt.Scan(&lamps[j])
		}
		sort.Slice(lamps, func(i, j int) bool { return lamps[i] < lamps[j] })
		for j := 1; j < N; j++ {
			spaces[j-1] = lamps[j] - lamps[j-1]
		}
		var max int
		for j := 0; j < N-1; j++ {
			if max < spaces[j] {
				max = spaces[j]
			}
		}
		fmt.Println(float64(max) / 2.0)
	}
}
