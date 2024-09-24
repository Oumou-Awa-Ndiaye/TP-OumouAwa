package main

import (
	"fmt"
)

func Principale() {
	// Test de la fonction Ft_missing
	fmt.Println(Ft_missing([]int{3, 1, 2}))                   // résultat : 0
	fmt.Println(Ft_missing([]int{0, 1}))                      // résultat : 2
	fmt.Println(Ft_missing([]int{9, 6, 4, 2, 3, 5, 7, 0, 1})) // résultat : 8
}
func Ft_missing(nums []int) int {
	n := len(nums)

	SommeAttendue := n * (n + 1) / 2

	SommeActuel := 0
	for _, num := range nums {
		SommeActuel += num
	}
	return SommeAttendue - SommeActuel
}

