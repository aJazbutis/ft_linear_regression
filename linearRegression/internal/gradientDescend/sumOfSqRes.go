package gradientDescend

import "math"

func SumOfSqRes(data [][]float64, intercept, slope float64) float64 {
	var sum float64
	for i := range data[0] {
		sum += math.Pow((data[1][i] - (intercept + slope*data[0][1])), 2)
	}
	return sum
}

/*derivatives*/
func DerInteceptSumOfSqRes(data [][]float64, intercept, slope float64) float64 {
	var sum float64
	for i := range data[0] {
		sum += -2 * (data[1][i] - (intercept + slope*data[0][i]))
	}
	return sum
}

func DerSlopeSumOfSqRes(data [][]float64, intercept, slope float64) float64 {
	var sum float64
	for i := range data[0] {
		sum += -2 * data[0][i] * (data[1][i] - (intercept + slope*data[0][i]))
	}
	return sum
}
