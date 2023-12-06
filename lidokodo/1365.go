package main

func smallerNumbersThanCurrent(nums []int) []int {

	count := make([]int, 101)
	re := make([]int, len(nums))

	for _, v := range nums {
		count[v] += 1
	}

	for i, v := range nums {
		summed := 0
		for _, w := range count[0:v] {
			summed += w
		}
		re[i] = summed
	}

	return re
}
