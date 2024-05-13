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

func PutPoints(points [][]float64) {
	dimensions := 2
	persist := false
	debug := false
	plot, _ := glot.NewPlot(dimensions, persist, debug)
	// fmt.Println(plot)
	
	points = Transpose(points)
	fmt.Println(points)
	// plot.AddPointGroup(name, style, points)
	style := "lines"
	// mean := Mean(points[1])
	// n, x := MinMax(points[0])
	// fmt.Println(mean, n, x)
	// name = "average"
	b0, b1 := MinSq(points)
	// // fmt.Println(R2(points, b0, b1))
	name := "minsqr"
	pred := line(b0, b1, points[0])
	// plot.AddPointGroup(name, style, [][]float64{{n, x}, {mean, mean}})
	// name = "regr"
	plot.AddPointGroup(name, style, [][]float64{points[0], pred})
	// theta := GradientDescent(points)
	// name = "gradient"
	// pred = line(theta[0], theta[1], points[0])
	// fmt.Println(pred)
	// fmt.Println(theta)
	// pred := test(theta, points)
	// name = "Test"
	// style = "points"
	// pred = test([]float64{b0, b1}, points)
	// fmt.Println(pred)
	// plot.AddPointGroup(name, style, [][]float64{points[0], pred})
	fmt.Println(plot.SavePlot("testG.png"))
}