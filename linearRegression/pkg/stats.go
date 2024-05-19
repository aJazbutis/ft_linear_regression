import (
	"encoding/csv"
	"slices"
	// "fmt"
	"math"
	"os"
	"strconv"
)

func toFloat(data [][]string) [][]float64 {
	ret := make([][]float64, len(data))
	for i := range data {
		x, err := strconv.ParseFloat(data[i][0], 64)
		if err != nil {
			panic(err)
		}
		y, err := strconv.ParseFloat(data[i][1], 64)
		if err != nil {
			panic(err)
		}
		ret[i] = []float64{x, y}
	}
	return ret
}

func Transpose[T any](data [][]T) [][]T {
	ret := make([][]T, len(data[0]))
	for i := range ret {
		ret[i] = make([]T, len(data))
	}
	for i := range data {
		for j := range data[i] {
			ret[j][i] = data[i][j]
		}
	}
	return ret
}

func Mean(data []float64) float64 {
	var sum float64
	for _, v := range data {
		sum += v
	}
	return sum / float64(len(data))
}

/*
SS(mean) sum of squares around mean
*/
func SsMean(data []float64, mean float64) float64 {
	var ss float64
	for _, v := range data {
		ss += math.Pow((v - mean), 2)
	}
	return ss
}

/*
variation: SS/n
*/
func VarMean(data []float64, mean float64) float64 {
	return SsMean(data, mean) / float64(len(data))
}

func StdDev(data []float64, mean float64) float64 {
	return math.Sqrt(VarMean(data, mean))
}

func SsFit(data [][]float64, b0, b1 float64) float64 {
	var ss float64
	for i := range data[0] {
		ss += math.Pow((data[1][i] - (b0 + b1*data[0][i])), 2)
	}
	return ss
}

func VarFit(data [][]float64, b0, b1 float64) float64 {
	return SsFit(data, b0, b1) / float64(len(data[0]))
}

func R2(data [][]float64, b0, b1 float64) float64 {
	varMean := VarMean(data[1], Mean(data[1]))
	varFit := VarFit(data, b0, b1)
	return (varMean - varFit) / varMean
}

func MinMax(data []float64) (float64, float64) {
	n, x := data[0], data[0]
	for _, v := range data {
		n, x = min(n, v), max(x, v)
	}
	return n, x
}

func quantile(data []float64, m, q int) float64 {
	n := len(data)
	var i int
	var ret float64
	i = (m*n)/q
	if n%q == 1 {
		i++
		ret = data[i]
	} else {
		ret = (data[i]+data[i+1])/2
	}
	return ret
}

func Quantiles(data []float64) (float64, float64, float64) {
	d := make([]float64, len(data))
	copy(d, data)
	slices.Sort(d)
	var q1, q2, q3 float64
	q1 = quantile(d, 1, 4)
	/*median*/
	q2 = quantile(d, 1, 2)
	q3 = quantile(d, 3, 4)
	return q1, q2, q3
}

func ParseCsv(path string) [][]float64 {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	reader := csv.NewReader(f)
	result, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}
	// for _, v := range result {
		// fmt.Println(v)
	// }
	ret := toFloat(result[1:])
	// ret = Transpose(ret)
	// fmt.Println(ret)
	return ret
}
