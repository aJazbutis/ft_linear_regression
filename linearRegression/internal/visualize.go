package internal

import (
	"fmt"
	"github.com/arafatk/glot"
)

func line(b0, b1 float64, data []float64) []float64 {
	ret := make([]float64, len(data))
	for i := range data {
		ret[i] = b0 + b1 * data[i]
	} 
	return ret
}

func PutPoints(points [][]float64) {
	dimensions := 2
	persist := false
	debug := false
	plot, _ := glot.NewPlot(dimensions, persist, debug)
	name := "Test"
	style := "points"
	points = Transpose(points)
	fmt.Println(points)
	plot.AddPointGroup(name, style, points)
	style = "lines"
	mean := Mean(points[1])
	n, x := MinMax(points[0])
	fmt.Println(mean, n, x)
	name = "average"
	b0, b1 := MinSq(points)
	fmt.Println(R2(points, b0, b1))
	pred := line(b0, b1, points[0])
	plot.AddPointGroup(name, style, [][]float64{{n, x}, {mean, mean}})
	name = "regr"
	plot.AddPointGroup(name, style, [][]float64{points[0], pred})
	plot.SavePlot("test.png")
}