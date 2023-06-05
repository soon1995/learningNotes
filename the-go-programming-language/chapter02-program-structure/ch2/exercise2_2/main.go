// Exercise 2.2: Write a general-purpose unit-conversion program analogous to cf that
// reads numbers from its command-line arguments or from the standard input if there are
// no arguments, and converts each number into units like temperature in Celsius and
// Fahrenheit,length in feet and meters, weight in pounds and kilograms, and the like.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"example.go/ch2/exercise2_2/length"
	"example.go/ch2/exercise2_2/temperature"
	"example.go/ch2/exercise2_2/weight"
)

func main() {
	if len(os.Args[1:]) == 0 {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			line := input.Text()
			cf(line)
		}
	} else {
		for _, arg := range os.Args[1:] {
			cf(arg)
		}
	}
}

func cf(arg string) {
	t, err := strconv.ParseFloat(arg, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cf:  %v\n", err)
		os.Exit(1)
	}
	f := temperature.Fahrenheit(t)
	c := temperature.Celsius(t)

	pound := weight.Pound(t)
	kg := weight.Kilogram(t)

	feet := length.Feet(t)
	meter := length.Meter(t)

	fmt.Printf("%s = %s, \t%s = %s\t, %s = %s\t, %s = %s\t, %s = %s\t, %s = %s\n",
		f, temperature.FtoC(f), c, temperature.CToF(c),
		pound, weight.PoundToKilogram(pound), kg, weight.KilogramToPound(kg),
		feet, length.FeetToMeter(feet), meter, length.MeterToFeet(meter))

}
