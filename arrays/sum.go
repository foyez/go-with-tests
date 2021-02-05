package arrays

// Sum takes an array and returns total
func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

// SumAllTails takes number of slices and return a slice
func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:] // slice[low:high]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}