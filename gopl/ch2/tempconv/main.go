package main

import (
	"flag"
	"fmt"
)

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func main() {
	fmt.Println("Temp Celsius/Fahrenheit")

	temp := flag.Float64("temp", 0, "temp")
	typeOfTemp := flag.String("type", "", "type of temp")

	flag.Parse()

	if *typeOfTemp == "C" {
		fmt.Println(CToF(Celsius(*temp)))
	} else if *typeOfTemp == "F" {
		fmt.Println(FToC(Fahrenheit(*temp)))
	} else {
		fmt.Println("Type can be C or F")
	}

}
