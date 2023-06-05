// Add Support for Kelvin temperatures to tempflag
package main

import (
	"flag"
	"fmt"
	"time"
)

// *celsius Flag satisfies the flag.Value interface
type celsiusFlag struct{ Celsius }

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit) // no error check needed
	switch unit {
	case "C", "°C":
		f.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = FToC(Fahrenheit(value))
		return nil
	case "K":
		f.Celsius = KToC(Kelvin(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

type Celsius float64

type Fahrenheit float64

type Kelvin float64

func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }

const (
	AbsoulateZeroC Celsius = -273.15
	FreezingC      Celsius = 0
	BoilingC       Celsius = 100
)

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func KToC(k Kelvin) Celsius {
	return Celsius(k - 273.15)
}

// CelsiusFlag defines a Celsius flag with the specified name,
// default value, and usage, and requtns the address of the flag variable.
// The flag argument must have a quantity and a unit, e.g., 100C,
func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}

var period = flag.Duration("period", 1*time.Second, "sleep period")

var temp = CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
	// var c ByteCounter
	// c.Write([]byte("hello"))
	// fmt.Println(c) // "5" , = len("hello")

	// c = 0 // reset the counter
	// var name = "Dolly"
	// fmt.Fprintf(&c, "hello, %s", name)
	// fmt.Println(c) // "12", = len("hello, Dolly")

	// flag.Parse()
	// fmt.Printf("Sleeping for %v...", *period)
	// time.Sleep(*period)
	// fmt.Println("")

}
