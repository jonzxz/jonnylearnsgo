package main

func numberOfEmployeesWhoMetTarget(hours []int, target int) int {

	if len(hours) == 0 {
		return 0
	}

	count := 0

	for _, e := range hours {
		if e >= target {
			count += 1
		}
	}

	return count
}
