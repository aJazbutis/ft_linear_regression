package gradientDescend

func estimatePrice(theta []float64, mileage float64) float64 {
	return theta[0] + theta[1]*mileage
}
