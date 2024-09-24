package main

func Ft_missing(nums []int) int {
	n := len(nums)

	SommeAttendue := n * (n + 1) / 2

	SommeActuel := 0
	for _, num := range nums {
		SommeActuel += num
	}
	return SommeAttendue - SommeActuel
}
