package main

func runningSum(nums []int) []int {
	numsNew := make([]int, len(nums))

	sum := 0

	for i := range nums {
		sum += nums[i]
		numsNew[i] = sum
	}
	return numsNew
}
