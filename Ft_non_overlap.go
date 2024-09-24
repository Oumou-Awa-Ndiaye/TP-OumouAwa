package main

import (
	"fmt"
	"sort"
)

func main3() {
	fmt.Println(Ft_non_overlap([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}})) // résultat : 1
	fmt.Println(Ft_non_overlap([][]int{{1, 2}, {2, 3}}))                 // résultat : 0
	fmt.Println(Ft_non_overlap([][]int{{1, 2}, {1, 2}, {1, 2}}))         // résultat : 2
}

func Ft_non_overlap(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	count := 0
	end := intervals[0][1] // Fin du premier intervalle

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < end {
			count++
		} else {
			end = intervals[i][1]
		}
	}

	return count
}

