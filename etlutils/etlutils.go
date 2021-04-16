package etlutils

import "math"

func Mean(list []float64) float64 {
	var total float64
	for _, value := range list {
		total += value
	}

	return total / float64(len(list))
}

func Variance(list []float64) float64 {
	// Subtract the mean and square the result
	mean := Mean(list)
	sumOfMean := 0
	for _, value := range list {
		meanDiff := float64(math.Pow((value - mean), 2))
		sumOfMean += int(meanDiff)
	}
	return float64(sumOfMean) / float64(len(list))
}

func MinMaxNormalization(list []float64) []float64 {
	min := 99999999999.0
	for _, value := range list {
		if value < min {
			min = value
		}
	}
	max := 0.0
	for _, value := range list {
		if value > max {
			max = value
		}
	}

	var normalizedList []float64
	// Normalize the entire list
	for _, value := range list {
		normalizedList = append(normalizedList, (value-min)/(max-min))
	}

	return normalizedList
}
