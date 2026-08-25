package main

import (
	"flag"
	"fmt"
)

type Kilometer float64
type Meter float64

func KToM(k Kilometer) Meter { return Meter(k * 1000) }

func MToK(m Meter) Kilometer { return Kilometer(m / 1000) }

func main() {
	fmt.Println("---Conversion---")
	val := flag.Float64("val", 0, "value of the metric")
	typeOfMetric := flag.String("type", "km", "type of the metric")

	flag.Parse()

	if *typeOfMetric == "km" {
		fmt.Println(KToM(Kilometer(*val)))
	} else if *typeOfMetric == "m" {
		fmt.Println(MToK(Meter(*val)))
	} else {
		fmt.Println("type should be km or m")
	}
}
