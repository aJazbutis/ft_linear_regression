package gradientDescend

func TmpT0(data [][]float64, learningRate float64, theta []float64) float64 {
	var sum float64
	for i := range data[0] {
		sum += estimatePrice(theta, data[0][i]) - data[1][i]
	}
	return learningRate * (sum / float64(len(data[0])))
}

func TmpT1(data [][]float64, learningRate float64, theta []float64) float64 {
	var sum float64
	for i := range data[0] {
		sum += data[0][i] * (estimatePrice(theta, data[0][i]) - data[1][i])
	}
	return learningRate * (sum / float64(len(data[0])))
}
