package main

import (
	"math"
)

func differenceOfSum(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	sumOfIndividuals := 0
	sumOfAll := 0

	for _, num := range nums {
		sumOfAll += num
		for num > 0 {
			t := num % 10
			num = num / 10
			sumOfIndividuals += t
		}

	}

	return int(math.Abs(float64(sumOfAll) - math.Abs(float64(sumOfIndividuals))))
}
