package main

func numJewelsInStonesOld(jewels string, stones string) int {
	jewelsRune := []rune(jewels)
	stonesRune := []rune(stones)
	count := 0
	for _, v := range stonesRune {
		if contains(jewelsRune, v) {
			count += 1
		}
	}
	return count
}

func contains(runes []rune, single rune) bool {
	for _, v := range runes {
		if v == single {
			return true
		}
	}
	return false
}

func numJewelsInStones(jewels string, stones string) int {
	present := make(map[rune]bool)
	count := 0
	for _, jewel := range jewels {
		present[jewel] = true
	}

	for _, stone := range stones {
		if _, ok := present[stone]; ok {
			count++
		}
	}
	return count

}
