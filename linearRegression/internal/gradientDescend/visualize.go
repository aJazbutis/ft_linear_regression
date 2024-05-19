package gradientDescend

import (
	"linearRegression/pkg/estimatePrice"

	"github.com/arafatk/glot"
)
const (
	dimensions = 2
	persist = false
	debug = false
)

func Scatter(data [][]float64, name, xAxis, yAxis, path string) {
	plot, err := glot.NewPlot(dimensions, persist, debug)
	if err != nil {
		panic(err)
	}
	defer plot.Close()
	style := "points"
	plot.AddPointGroup(name, style, data)
	plot.SetXLabel(xAxis)
	plot.SetYLabel(yAxis)
	plot.SavePlot(path)
}

func predict(data, theta []float64) []float64 {
	ret := make([]float64, len(data))
	for i := range data {
		ret[i] = estimatePrice.EstimatePrice(theta, data[i])
	}
	return ret
}

func Visualize(data Data) {
	plot, _ := glot.NewPlot(dimensions, persist, debug)
	defer plot.Close()
	/*
		csv data scatter
	*/
	name := "Data"
	style := "points"
	plot.AddPointGroup(name, style, data.Data)
	plot.SetXLabel("Mileage")
	plot.SetYLabel("Price")
	plot.SavePlot("Data.png")
	plot.ResetPlot()
	/*
		normalized dataset and linear regression
	*/
	name = "Normalised data"
	plot.SetLabels("x", "f(x)")
	plot.AddPointGroup(name, style, data.Norm)
	
	style = "lines"
	name = "Line of Regression"
	prediction := predict(data.Norm[0], data.Theta)
	plot.AddPointGroup(name, style, [][]float64{data.Norm[0], prediction})
	plot.SavePlot("NormalizedData.png")
	plot.ResetPlot()
	/*
		rescale	Thetas, scatter and regression
	*/
	data.rescaleThetas()
	name = "Data"
	style = "points"
	plot.SetLabels("Mileage", "Price")
	plot.AddPointGroup(name, style, data.Data)

	name = "Line of Regression"
	prediction = predict(data.Data[0], data.Theta)
	plot.AddPointGroup(name, style, [][]float64{data.Data[0], prediction})
	plot.SavePlot("Result.png")	
	// b0, b1 := MinSq(data)
	// fmt.Println(R2(data, b0, b1))
	// name = "minsqr"
	// pred := line(b0, b1, data[0])
	// plot.AddPointGroup(name, style, [][]float64{{n, x}, {mean, mean}})
	// name = "regr"
	// plot.AddPointGroup(name, style, [][]float64{data[0], pred})

	// fmt.Println(b0, b1)

	// pred := test(theta, data)
	// name = "Test"
	style = "lines"
	// pred = test([]float64{b0, b1}, data)
	// fmt.Println(pred)
	// plot.AddPointGroup(name, style, [][]float64{normData[0], pred})
	// plot.AddPointGroup(name, style, [][]float64{data[0], pred})
	// fmt.Println(plot.SavePlot("testG.png"))
	// plot.ResetPlot()
	// name = "data"
	// style = "points"
	// plot.AddPointGroup(name, style, data)
	// name = "data"
	// style = "points"
	// plot.AddPointGroup(name, style, normData)
	// name = "model"
	// style = "lines"
	// theta[1] = (maxY - minY) * theta[1] / (maxX - minX)
	// theta[0] = minY + ((maxY - minY) * theta[0]) + theta[1]*(1-minX)
	// fmt.Println(theta)
	// plot.AddPointGroup(name, style, [][]float64{data[0], pred})

	// plot.SavePlot("grad.png")
}
