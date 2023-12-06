package main

func shuffle(nums []int, n int) []int {
	var numsCpy []int

	for i := 0; i < n; i++ {
		numsCpy = append(numsCpy, nums[i], nums[i+n])
	}
	return numsCpy
}
