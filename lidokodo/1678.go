package main

func interpret(command string) string {
	keys := map[string]string{
		"G":    "G",
		"()":   "o",
		"(al)": "al",
	}
	var result string
	var substr string

	for _, c := range command {
		substr += string(c)

		if val, ok := keys[substr]; ok {
			result += val
			substr = ""
		}
	}

	return result
}
