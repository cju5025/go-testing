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

	for _, numbers := range numbersToSum {
		sums = append(sums, SumSlice(numbers))
	}

	return sums
}
