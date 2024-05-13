package internal

import "math"

func MinSq(data [][]float64) (float64, float64) {
	var b0, b1 float64
	xMean := Mean(data[0])
	yMean := Mean(data[1])
	var num, denum float64
	for i := range data[0] {
		num += (data[0][i] - xMean) *(data[1][i] - yMean)
		denum += math.Pow((data[0][i] - xMean), 2)
	}
	b1 = num/denum
	b0 = yMean - b1*xMean
	return b0, b1
}

func MinSqEstimate(b0, b1, x float64) float64 {
	return b0 + b1*x
}