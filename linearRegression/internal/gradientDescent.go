package internal

import (
	"fmt"
	"math"
	"math/rand"
)

const learningRate = 0.01

type Data struct {
	Data	[][]float64
	MeanX	float64
	MeanY	float64
	MinX	float64
	MaxX	float64
	MinY	float64
	MaxY	float64
}

func GradientDescent(data [][]float64) []float64 {
	// learningRate := 0.01
	iterations := 10000
	minStepSize := 0.00001
	theta := []float64{rand.Float64(), rand.Float64()}
	for it := 0; it < iterations; it++ {
		tmpT0, tmpT1 := TmpT0(data, learningRate, theta), TmpT1(data, learningRate, theta)
		stepSize := math.Sqrt(math.Pow(tmpT0, 2) + math.Pow(tmpT1, 2))
		if stepSize <= minStepSize {
			fmt.Println("Iterations: ", it+1)
			break
		}
		theta[0] -= tmpT0
		theta[1] -= tmpT1
	}
	return theta
}