package main

func maximumWealth(accounts [][]int) int {
	highest := 0

	for _, v := range accounts {
		total := getSum(v)

		if total > highest {
			highest = total
		}
	}
	return highest
}

func getSum(account []int) int {
	sum := 0
	for _, v := range account {
		sum += v
	}
	return sum
}
