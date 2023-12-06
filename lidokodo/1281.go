package main

func subtractProductAndSum(n int) int {

	sumOfNumber := 0
	productOfNumber := 1

	for n > 0 {
		t := n % 10
		sumOfNumber += t
		productOfNumber *= t
		n /= 10
	}

	return productOfNumber - sumOfNumber
}

func main() {
	subtractProductAndSum(234)
}
