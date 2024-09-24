package main

import (
	"fmt"
)

func Ft_coin(coins []int, amount int) int {

	result := solve(coins, amount, 0)

	if result == 1000000 {
		return -1
	}
	return result
}

//Fonction récursive pour trouver le nombre minimal de pièces

func solve(coins []int, amount int, count int) int {
	if amount == 0 {
		return count
	}

	if amount < 0 {
		return 1000000
	}

	minCoins := 1000000
	for _, coin := range coins {

		temp := solve(coins, amount-coin, count+1)
		if temp < minCoins {
			minCoins = temp
		}
	}
	return minCoins
}

func main() {
	fmt.Println(Ft_coin([]int{1, 2, 5}, 11)) // resultat : 3
	fmt.Println(Ft_coin([]int{2}, 3))        // resultat : -1
	fmt.Println(Ft_coin([]int{1}, 0))        // resultat : 0
}
