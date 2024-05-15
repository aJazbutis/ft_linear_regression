package internal

import (
	"fmt"
	"github.com/arafatk/glot"
	"math"
)

func line(b0, b1 float64, data []float64) []float64 {
	ret := make([]float64, len(data))
	for i := range data {
		ret[i] = b0 + b1 * data[i]
	} 
	return ret
}

func test(theta []float64, data [][]float64) [] float64 {
	ret := make([]float64, len(data[0]))
	for i := range ret {
		ret[i] = math.Pow((data[1][i] - (theta[0] + theta[1]*data[0][i])), 2)
	}
	return ret
}

func normalize(data [][]float64) [][]float64 {
	ret := make([][]float64, len(data))
	for i := range ret {
		ret[i] = make([]float64, len(data[i]))
	}
	for i := range data {
		mean := Mean(data[i])
		// varMean := VarMean(data[i], mean)
		stdDev := StdDev(data[i], mean)
		for j := range data[i] {
			ret[i][j] = (data[i][j] - mean) / stdDev
		}
	}
	return ret
}

func normalizeMinMax(data [][]float64) [][]float64 {
	ret := make([][]float64, len(data))
	for i := range ret {
		ret[i] = make([]float64, len(data[i]))
	}
	for i := range data {
		n, x := MinMax(data[i])
		for j := range data[i] {
			ret[i][j] = (data[i][j] - n)/(x-n)
		}
	}
	return ret
}

func Visualize(data [][]float64) {
	dimensions := 2
	persist := false
	debug := false
	plot, _ := glot.NewPlot(dimensions, persist, debug)
	defer plot.Close()
	// fmt.Println(plot)
	name := "data"
	style := "points"	
	data = Transpose(data)
	fmt.Println(data)
	plot.AddPointGroup(name, style, data)
	plot.SavePlot("DataScatter.png")
	// style = "lines"
	// mean := Mean(data[1])
	// n, x := MinMax(data[0])
	// fmt.Println(mean, n, x)
	// name = "average"
	b0, b1 := MinSq(data)
	// fmt.Println(R2(data, b0, b1))
	// name = "minsqr"
	// pred := line(b0, b1, data[0])
	// plot.AddPointGroup(name, style, [][]float64{{n, x}, {mean, mean}})
	// name = "regr"
	// plot.AddPointGroup(name, style, [][]float64{data[0], pred})
	plot.ResetPlot()
	// normData := normalize(data)	
	normData := normalizeMinMax(data)
	name = "NORM"
	style = "points"
	plot.AddPointGroup(name, style, normData)
	// fmt.Println(normData)
	theta := GradientDescent(normData)
	// pred := line(theta[0], theta[1], data[0])
	pred := line(theta[0], theta[1], normData[0])
	// for i := range pred {
		// fmt.Println(data[0][i], pred[i])
	// }
	// fmt.Println(pred)
	name = "gradient"
	// pred = line(theta[0], theta[1], data[0])
	// fmt.Println(pred)

	fmt.Println(b0, b1)
	fmt.Println(theta)
	
	// pred := test(theta, data)
	// name = "Test"
	style = "lines"
	// pred = test([]float64{b0, b1}, data)
	// fmt.Println(pred)
plot.AddPointGroup(name, style, [][]float64{normData[0], pred})
	// plot.AddPointGroup(name, style, [][]float64{data[0], pred})
	fmt.Println(plot.SavePlot("testG.png"))
	plot.ResetPlot()
	name = "data"
	style = "points"
	plot.AddPointGroup(name, style, data)
	// name = "data"
	// style = "points"
	// plot.AddPointGroup(name, style, normData)
	name = "model"
	style = "lines"
	minY, maxY := MinMax(data[1])
	minX, maxX := MinMax(data[0])
	// mean := Mean(data[0])
	// stdDev := StdDev(data[0], mean)
	// theta[0] = theta[0] * StdDev(data[1], Mean(data[1])) +Mean(data[1]) - (theta[1] * StdDev(data[0], Mean(data[0])) + Mean(data[1]))
	theta[1] = (maxY-minY)* theta[1]/(maxX-minX)
	theta[0] = minY + ((maxY-minY)*theta[0]) + theta[1]*(1-minX)
	fmt.Println(theta)
	pred = line(theta[0], theta[1], data[0])
	plot.AddPointGroup(name, style, [][]float64{data[0], pred})

	plot.SavePlot("grad.png")
}