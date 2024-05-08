package internal

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func toFloat(data [][]string) [][]float64 {
	ret := make([][]float64, len(data))
	for i := range data {
		x, err  := strconv.ParseFloat(data[i][0], 64)
		if err != nil {
			panic(err)
		}
		y, err := strconv.ParseFloat(data[0][1], 64)
		if err != nil {
			panic(err)
		}
		ret[i] = []float64{x, y}
	}
	return ret
}

func ParseCsv(path string) [][]float64 {
	f , err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	reader := csv.NewReader(f)
	result, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}
	for _, v := range result {
		fmt.Println(v)
	}
	ret := toFloat(result[1:])
	fmt.Println(ret)
	return ret
}