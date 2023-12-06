package main

func numIdenticalPairs(nums []int) int {
	occurences := make(map[int]int)
	sum := 0
	for _, v := range nums {
		_, exist := occurences[v]
		if !exist {
			occurences[v] = 1
		} else {
			sum += occurences[v]
			occurences[v] += 1
		}
	}

	return sum
}
