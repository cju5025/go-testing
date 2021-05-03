package main

func SumArray(numbers [5]int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumSlice(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	lengthOfNumbersToSum := len(numbersToSum)
	sums := make([]int, lengthOfNumbersToSum)

	for i, numbers := range numbersToSum {
		sums[i] = SumSlice(numbers)
	}

	return sums
}
