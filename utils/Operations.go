package utils

import (
	"math"
)

func Averge(numbers []float64) float64 {
	var sum, average float64

	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}
	average = math.Round(sum / float64(len(numbers)))
	return average
}

func Median(numbers []float64) float64 {
	var sum, median float64

	if len(numbers)%2 == 1 {
		median = numbers[len(numbers)/2]
	} else if len(numbers)%2 == 0 {
		sum = numbers[(len(numbers)/2)] + numbers[(len(numbers)/2)-1]
		median = math.Round(sum / 2)
	}
	return median
}

func Variance(nb []float64, average float64) float64 {
	var vSum float64

	for i := 0; i < len(nb); i++ {
		diff := nb[i] - average
		vSum += diff * diff
	}

	return math.Round(vSum / float64(len(nb)))
}
