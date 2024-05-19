package main

import (
	"flag"
	"linearRegression/internal/gradientDescend"
	"log"
)

func main()  {
	defer func ()  {
		if r := recover(); r != nil {
			log.Fatal(r)
		}
	}()
	f := flag.String("f", "../../../data.csv", "csv file, format: x, y")
	flag.Parse()
	gradientDescend.Descend(*f)
}