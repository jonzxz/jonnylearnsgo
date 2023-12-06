package main

func kidsWithCandies(candies []int, extraCandies int) []bool {
	maxCandies := 0
	for _, v := range candies {
		if v > maxCandies {
			maxCandies = v
		}
	}

	results := make([]bool, len(candies))
	for i, v := range candies {
		if v+extraCandies >= maxCandies {
			results[i] = true
		} else {
			results[i] = false
		}
	}

	return results
}
