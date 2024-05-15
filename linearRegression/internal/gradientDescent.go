package internal

import (
	"fmt"
	"math"
	"math/rand"
	// "fmt"
	// "math"
	// "math/rand"
)

func GradientDescent(data [][]float64) []float64 {
	learningRate := 0.01
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
		// fmt.Println(theta, stepSize)
	}
	return theta
	// intercept := rand.Float64()
	// slope := rand.Float64()
	// stepNum := 1000
	// learningRate := 0.001
	// theta := []float64{0, 0}
	// minStepSize := 0.001
	// _, m := MinMax(data[1])
	// theta := []float64{float64(m), -rand.Float64()}
	// learningRate := 0.01
	// maxStepNum := 1000
	// fmt.Println(theta, "QQ")
	// slope0 := DerInteceptSumOfSqRes(data, theta[0], theta[1])
	// slope1 := DerSlopeSumOfSqRes(data, theta[0], theta[1])
	// stepSize0 := slope0 * learningRate
	// stepSize1 := slope1 * learningRate
	// theta[0], theta[1] = theta[0]-stepSize0, theta[1]-stepSize1
	// for (stepSize0 > minStepSize || stepSize1 > minStepSize) && maxStepNum > 0 {
		// slope0 := DerInteceptSumOfSqRes(data, theta[0], theta[1])
	// 	slope1 := DerSlopeSumOfSqRes(data, theta[0], theta[1])
	// 	stepSize0 := slope0 * learningRate
	// 	stepSize1 := slope1 * learningRate
	// 	theta[0], theta[1] = theta[0]-stepSize0, theta[1]-stepSize1	
	// 	fmt.Print(theta)
	// 	maxStepNum--
	// }
	// fmt.Println()
	
/*
	tmpT0 := TmpT0(data, learningRate, theta)
	tmpT1 :=  TmpT1(data, learningRate, theta)
	theta[0], theta[1] = theta[0] - tmpT0, theta[1] - tmpT1
	for (tmpT0 > minStepSize || tmpT1 > minStepSize) && maxStepNum > 0 {
		if math.Abs(tmpT0) > minStepSize {
			tmpT0 = TmpT0(data, learningRate, theta)
		}
		if math.Abs(tmpT1) > minStepSize {
			tmpT1 = TmpT1(data, learningRate, theta)	
		}
		theta[0], theta[1] = theta[0] - tmpT0, theta[1] - tmpT1
	}*/
	// return theta

}