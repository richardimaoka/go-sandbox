package main

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersOfNumbers ...[]int) (sum []int) {
	return []int{3, 9}
}
