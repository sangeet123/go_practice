package main

import (
	"./tempconv"
	"fmt"
	"os"
	"strconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}

		// type conversion from float64 to Celcious and Farenheit
		c := tempconv.Celsius(t)
		f := tempconv.Farenheit(t)

		fmt.Printf("%s=%s, %s=%s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
	}
}
