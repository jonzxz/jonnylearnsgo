package main

func buildArray(nums []int) []int {
	numsNew := make([]int, len(nums))

	for i := range nums {
		numsNew[i] = nums[nums[i]]
	}
	return numsNew
}
