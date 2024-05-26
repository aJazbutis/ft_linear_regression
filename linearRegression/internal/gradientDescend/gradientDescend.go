package gradientDescend

import (
	"fmt"
	"linearRegression/pkg/stats"
	"math"
	"math/rand"
)

const (
	learningRate	= 0.01
	iterations		= 10000
	minStepSize		= 0.00001
)

type Data struct {
	Data	[][]float64
	Norm	[][]float64
	MeanX	float64
	MeanY	float64
	MinX	float64
	MaxX	float64
	MinY	float64
	MaxY	float64
	Theta	[]float64
}

/*
	scale back to fit orginal data
*/
func (data *Data) rescaleThetas() {
	data.Theta[1] = (data.MaxY - data.MinY) * data.Theta[1] / (data.MaxX - data.MinX)
	data.Theta[0] = data.MinY + ((data.MaxY-data.MinY) * data.Theta[0]) + data.Theta[1]*(1-data.MinX)
}

func getData(path string) Data{
	var data Data
	data.Data = stats.Transpose(stats.ParseCsv(path))
	data.MinX, data.MaxX = stats.MinMax(data.Data[0])
	data.MinY, data.MaxY = stats.MinMax(data.Data[1])
	data.Norm = stats.Normalize(data.Data)
	return data
}

func Descend(path string)  {
	data := getData(path)
	//VisualizeData(&data)

	// descendSumOfSqRes(&data)
	descendAverageRes(&data)
	prediction := predict(data.Norm[0], data.Theta)
	mse := stats.MSE([][]float64{data.Norm[1], prediction})
	mae := stats.MAE([][]float64{data.Norm[1], prediction})
	//VisualizeNormData(&data, prediction)
	fmt.Println("MSE:", mse)
	fmt.Println("MAE:", mae)
	data.rescaleThetas()
	prediction = predict(data.Data[0], data.Theta)
	//VisualizeRescaled(&data, prediction)
	r2 := stats.R2(data.Data, data.Theta[0], data.Theta[1])
	fmt.Println("r^2:",  r2)
	fmt.Println("theta0:", data.Theta[0])
	fmt.Println("theta1:", data.Theta[1])
}

/* better one */
func descendSumOfSqRes(data *Data) {
	data.Theta = []float64{rand.Float64(), rand.Float64()}
	for it := 0; it < iterations; it++ {
		intercept := DerInteceptSumOfSqRes(data.Norm, data.Theta[0], data.Theta[1])
		slope := DerSlopeSumOfSqRes(data.Norm, data.Theta[0], data.Theta[1])
		stepSizeIntercept := intercept * learningRate
		stepSizeSlope := slope * learningRate
		stepSize := math.Sqrt(math.Pow(stepSizeIntercept, 2) + math.Pow(stepSizeSlope, 2))
		if stepSize <= minStepSize {
			fmt.Println("Iterations: ", it+1)
			break
		}
		data.Theta[0] -= stepSizeIntercept
		data.Theta[1] -= stepSizeSlope

	}
}


/* provided formulas */
func descendAverageRes(data *Data) {
	data.Theta = []float64{rand.Float64(), rand.Float64()}
	for it := 0; it < iterations; it++ {
		tmpT0, tmpT1 := TmpT0(data.Norm, learningRate, data.Theta), TmpT1(data.Norm, learningRate, data.Theta)
		stepSize := math.Sqrt(math.Pow(tmpT0, 2) + math.Pow(tmpT1, 2))
		if stepSize <= minStepSize {
			fmt.Println("Iterations: ", it+1)
			break
		}
		data.Theta[0] -= tmpT0
		data.Theta[1] -= tmpT1
	}
}
