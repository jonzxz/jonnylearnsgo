package main

func getConcatenation(nums []int) []int {
	numsNew := make([]int, len(nums)*2)

	for i := 0; i < len(nums); i++ {
		numsNew[i] = nums[i]
		numsNew[i+len(nums)] = nums[i]
	}

	return numsNew
}
