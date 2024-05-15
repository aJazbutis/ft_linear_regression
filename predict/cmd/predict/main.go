package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
)

func main() {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println(r)
		}
	}()
	t0 := flag.Float64("t0", 0.0, "theta0")
	t1 := flag.Float64("t1", 0.0, "theta1")
	flag.Parse()
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Input mileage:")
	scanner.Scan()
	mileage, err := strconv.ParseFloat(scanner.Text(), 64)
	if err != nil {
		panic("Invalid mileage input")
	}
	if mileage < 0.0 {
		panic("Negative mileage")
	}
	fmt.Println("Predicted price: ", *t0+*t1*mileage)
}
