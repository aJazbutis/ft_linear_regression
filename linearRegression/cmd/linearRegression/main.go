package main

import (
	"flag"
	"log"
	"gradient/internal"
)

func main()  {
	defer func ()  {
		if r := recover(); r != nil {
			log.Fatal(r)
		}
	}()
	f := flag.String("f", "../../../data.csv", "csv file, format: x, y")
	flag.Parse()
	internal.Visualize(internal.ParseCsv(*f))
}