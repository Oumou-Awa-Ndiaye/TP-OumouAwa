package main

func Ft_coin(coins []int, amount int) int {
	//D'abord ici il n'y aura pas de pièce si le montant est 0
	if amount == 0 {
		return 0
	}

	// Là j'ai inicialisé juste une valeur pour stocker le nombre minimal de pièces
	minCoins := -1

	// Là j'ai créée ma boucle pour parcourir toutes les pièces
	for _, coin := range coins {
		if coin <= amount {
			// Je calcule combien de pièces de cette valeur sont nécessaires pour atteindre amount
			temp := Ft_coin(coins, amount-coin)
			if temp != -1 {
				temp += 1

				if minCoins == -1 || temp < minCoins {
					minCoins = temp
				}
			}
		}
	}

	return minCoins
}
