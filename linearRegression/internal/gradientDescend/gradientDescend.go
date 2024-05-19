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
	descend(&data)
	
	// Visualize(data)
	fmt.Println(data.Theta)
	data.rescaleThetas()
	fmt.Println(data.Theta)
}

func descend(data *Data) {
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
/*
func descend(data [][]float64) []float64 {
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
*/