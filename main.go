package math

func Add(numbers ...int) int {
	sum := 0

	for _, num := range numbers {
		sum += num
	}

	return sum
}

func Subtract(numbers ...int) int {
	sum := 0

	for _, num := range numbers {
		sum -= num
	}

	return sum
}

func Divide(a, b int) int {
	return a / b
}
