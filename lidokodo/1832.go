package main

func checkIfPangram(sentence string) bool {
	mapping := make(map[string]int)

	for i := 'a'; i <= 'z'; i++ {
		t := string(i)
		mapping[t] = 0
	}

	for _, c := range sentence {
		t := string(c)
		mapping[t] += 1
	}

	for _, v := range mapping {
		if v == 0 {
			return false
		}

	}

	return true
}
